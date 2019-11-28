package echoT

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var RouterMap = make(map[string]*RouterInfo, 0)

type InterfaceMap map[string]reflect.Type
type RouterDoc struct {
	Title string
	Desc  string
}
type RouterInfo struct {
	Doc          RouterDoc
	Mapping      string
	HttpMethod   string
	Auth         bool
	Do           interface{}
	f            *Func
	rout         bool
	InterfaceMap InterfaceMap
}

func (r *RouterInfo) DoAuth() RouterInfo {
	r.Auth = true
	return *r
}

func R(r RouterInfo) RouterInfo {
	r.Router()
	return r
}
func GetMapping(Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodGet, Do: Do}
	return r
}
func PostMapping(Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodPost, Do: Do}
	return r
}
func DeleteMapping(Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodDelete, Do: Do}
	return r
}
func PutMapping(Mapping string, Do interface{}) *RouterInfo {
	r := &RouterInfo{Mapping: Mapping, HttpMethod: http.MethodPut, Do: Do}
	return r
}

func (r *RouterInfo) Router() RouterInfo {
	if r.rout {
		return *r
	}
	r.Mapping = strings.TrimLeft(r.Mapping, "/")
	if _, ok := RouterMap[r.Mapping+r.HttpMethod]; ok {
		panic(fmt.Sprintf("重复的路由配置 `/%s.%s`", r.Mapping, r.HttpMethod))
	}
	r.f = NewFunc(r.Do)
	RouterMap[r.Mapping+r.HttpMethod] = r
	r.rout = true
	han := func(ctx echo.Context) error {
		ses, e := session.Get("sessionId", ctx)
		if e != nil {
			return e
		}
		if ses.IsNew {
			ses.Options = &sessions.Options{
				Path:   "/",
				MaxAge: 60 * 60, // 以秒为单位
			}
		}

		if err := ses.Save(ctx.Request(), ctx.Response()); err != nil {
			return err
		}
		if out, err := r.f.Call(ctx); err != nil {
			return err
		} else if err := ctx.JSON(200, out); err != nil {
			return err
		}
		return nil
	}
	var m []echo.MiddlewareFunc
	if r.Auth {
		m = append(m, AuthMiddleware)
	}
	switch r.HttpMethod {
	case http.MethodGet:
		g.GET(r.Mapping, han, m...)
	case http.MethodPost:
		g.POST(r.Mapping, han, m...)
	case http.MethodPut:
		g.PUT(r.Mapping, han, m...)
	case http.MethodDelete:
		g.DELETE(r.Mapping, han, m...)
	}
	return *r
}
