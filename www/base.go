package www

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ltto/T/www/rest"
)

var (
	Engine *gin.Engine
)

func init() {
	Engine = gin.New()
	store := cookie.NewStore([]byte("secret"))
	Engine.Use(sessions.Sessions(SessionID, store))
}

type RouterInfo struct {
	Mapping    string
	HttpMethod string
	Routes     gin.IRoutes
	Do         interface{}
	f          *Func
}

func GetMapping(routes gin.IRoutes, Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodGet, Do: Do, Routes: routes}
	Router(r)
	return r
}
func PostMapping(routes gin.IRoutes, Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodPost, Do: Do, Routes: routes}
	Router(r)
	return r
}
func DeleteMapping(routes gin.IRoutes, Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodDelete, Do: Do, Routes: routes}
	Router(r)
	return r
}
func PutMapping(routes gin.IRoutes, Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodPut, Do: Do, Routes: routes}
	Router(r)
	return r
}
func Get(Mapping string, Do interface{}) *RouterInfo {
	return GetMapping(nil, Mapping, Do)
}
func Post(Mapping string, Do interface{}) *RouterInfo {
	return PostMapping(nil, Mapping, Do)
}
func Delete(Mapping string, Do interface{}) *RouterInfo {
	return DeleteMapping(nil, Mapping, Do)
}
func Put(routes gin.IRoutes, Mapping string, Do interface{}) *RouterInfo {
	return PutMapping(routes, Mapping, Do)
}

func Router(info *RouterInfo) {
	pathMap := path.Clean(info.Mapping)
	if info.Routes == nil {
		info.Routes = Engine
	}
	switch info.HttpMethod {
	case http.MethodGet:
		info.Routes.GET(pathMap, info.ginFunc())
	case http.MethodPost:
		info.Routes.POST(pathMap, info.ginFunc())
	case http.MethodPut:
		info.Routes.PUT(pathMap, info.ginFunc())
	case http.MethodDelete:
		info.Routes.DELETE(pathMap, info.ginFunc())
	}
}

func (info *RouterInfo) ginFunc() func(c *gin.Context) {
	info.f = NewFunc(info.Do)
	return func(c *gin.Context) {
		var (
			session = sessions.Default(c)
			ctx     = &Context{Context: c, Session: session}
			do      interface{}
			err     error
		)
		session.Options(sessions.Options{MaxAge: 60 * 60})

		c.Writer.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Add("Pragma", "no-cache")
		c.Writer.Header().Add("Expires", "0")
		c.Set("routerInfo", info)

		if do, err = info.f.Call(ctx); err != nil {
			DefaultErrHandler(ctx, err)
		}
		if err := session.Save(); err != nil {
			DefaultErrHandler(ctx, err)
			return
		}
		switch t := do.(type) {
		case *os.File:
			defer t.Close()
			c.Writer.Header().Add("Content-Disposition", "attachment; filename="+filepath.Base(t.Name()))
			ext := filepath.Ext(t.Name())
			c.Writer.Header().Add("Content-Type", ContentType(ext))
			if _, err := io.Copy(c.Writer, t); err != nil {
				DefaultErrHandler(ctx, err)
				return
			}
		case string:
			stringHdl(ctx, t)
		case *string:
			s := *t
			stringHdl(ctx, s)
		case io.ReadCloser:
			defer t.Close()
			if _, err := io.Copy(c.Writer, t); err != nil {
				DefaultErrHandler(ctx, err)
				return
			}
		case io.Reader:
			if _, err := io.Copy(c.Writer, t); err != nil {
				DefaultErrHandler(ctx, err)
				return
			}
		case error:
			DefaultErrHandler(ctx, t)
			return
		case nil:
		default:
			c.JSON(http.StatusOK, do)
		}
	}
}

func stringHdl(c *Context, s string) {
	if ss, ok := rest.Redirect(s); ok {
		c.Redirect(http.StatusMovedPermanently, ss)
		return
	}
	var ok bool
	if s, ok = rest.File(s); ok {
		contentType := ContentType(filepath.Ext(s))
		c.Writer.Header().Add("Content-Type", contentType)
		if strings.Contains(contentType, "image") {
			c.Writer.Header().Add("Content-Disposition", "filename="+filepath.Base(s))
		} else {
			c.Writer.Header().Add("Content-Disposition", "attachment;filename="+filepath.Base(s))
		}
		http.ServeFile(c.Writer, c.Request, s)
		return
	}
	if s, ok = rest.Html(s); ok {
		get := c.CParams()
		c.HTML(http.StatusOK, s, get)
		return
	}
	if strings.ToLower(path.Ext(s)) == ".html" {
		get := c.CParams()
		c.HTML(http.StatusOK, s, get)
		return
	}
	{
		paramsMap := c.CParams()
		params := make([]interface{}, 0, len(paramsMap))
		for k := range paramsMap {
			params = append(params, paramsMap[k])
		}
		c.String(http.StatusOK, s, params...)
	}
}
