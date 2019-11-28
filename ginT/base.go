package ginT

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/ltto/T/ginT/ctx"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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
	auth       bool
	Routes     gin.IRoutes
	Do         interface{}
}

func (r *RouterInfo) SetAuth(b bool) {
	r.auth = b
}

func GetMapping(Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodGet, Do: Do}
	Router(r)
	return r
}
func PostMapping(Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodPost, Do: Do}
	Router(r)
	return r
}
func DeleteMapping(Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodDelete, Do: Do}
	Router(r)
	return r
}
func PutMapping(Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodPut, Do: Do}
	Router(r)
	return r
}

func Router(info *RouterInfo) {
	pathMap := path.Clean(info.Mapping)
	if info.Routes == nil {
		info.Routes = Engine
	}
	switch info.HttpMethod {
	case http.MethodGet:
		info.Routes.GET(pathMap, ginFunc(info))
	case http.MethodPost:
		info.Routes.POST(pathMap, ginFunc(info))
	case http.MethodPut:
		info.Routes.PUT(pathMap, ginFunc(info))
	case http.MethodDelete:
		info.Routes.DELETE(pathMap, ginFunc(info))
	}
}

func ginFunc(info *RouterInfo) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Add("Pragma", "no-cache")
		c.Writer.Header().Add("Expires", "0")
		c.Set("routerInfo", info)
		context := &ctx.Context{Context: c}
		if interErr := DoInterceptorList(context); interErr.isErr {
			http.Error(c.Writer, interErr.Error(), interErr.Code)
		}
		if info.auth && !AuthFuc(context) {
			c.String(http.StatusUnauthorized, "unauthorized!")
			return
		}
		newFunc := NewFunc(reflect.ValueOf(info.Do))
		do := newFunc.Call(context)
		switch t := do.(type) {
		case *os.File:
			defer t.Close()
			c.Writer.Header().Add("Content-Disposition", "attachment; filename="+filepath.Base(t.Name()))
			ext := filepath.Ext(t.Name())
			switch ext {
			case ".jpg", ".JPG",
				".jpe", ".JPE",
				".jpeg", ".JPEG",
				".jfif", ".JFIF":
				c.Writer.Header().Add("Content-Type", "image/jpeg")
			case ".gif", ".GIF":
				c.Writer.Header().Add("Content-Type", "image/gif")
			case ".png", ".PNG":
				c.Writer.Header().Add("Content-Type", "image/png")
			case ".fax", ".FAX":
				c.Writer.Header().Add("Content-Type", "image/fax")
			case ".tif", ".TIF", ".tiff", ".TIFF":
				c.Writer.Header().Add("Content-Type", "image/tiff")
			case ".ico", ".ICO":
				c.Writer.Header().Add("Content-Type", "image/x-icon")
			case ".net", ".NET":
				c.Writer.Header().Add("Content-Type", "image/pnetvue")
			case ".rp", ".RP":
				c.Writer.Header().Add("Content-Type", "image/vnd.rn-realpix")
			case ".wbmp", ".WBMP":
				c.Writer.Header().Add("Content-Type", "image/vnd.wap.wbmp")
			default:
				c.Writer.Header().Add("Content-Type", "application/octet-stream")
			}
			_, err := io.Copy(c.Writer, t)
			if err != nil {
				panic(err)
			}
		case string:
			stringHdl(c, t)
		case *string:
			s := *t
			stringHdl(c, s)
		case error:
			c.String(http.StatusInternalServerError, t.Error())
		default:
			c.JSON(http.StatusOK, do)
		}
	}
}

func stringHdl(c *gin.Context, s string) {
	if strings.HasPrefix(s, HTTP_REDIRECT) {
		c.Redirect(http.StatusMovedPermanently, s[len(HTTP_REDIRECT)-1:])
		return
	} else if strings.HasPrefix(s, HTTP_FILE) {
		s = s[len(HTTP_FILE)-1:]
		c.Writer.Header().Add("Content-Disposition", "attachment; filename="+filepath.Base(s))
		c.Writer.Header().Add("Content-Type", "application/octet-stream")
		c.File(s)
		return
	} else if strings.HasPrefix(s, HTTP_IMG) {
		s = s[len(HTTP_FILE)-1:]
		c.Writer.Header().Add("Content-Type", "image/jpeg")
		open, _ := os.Open(s)
		_, _ = io.Copy(c.Writer, open)
		return
	}
	get, _ := c.Get("html")
	c.HTML(http.StatusOK, s+".html", get)
}
