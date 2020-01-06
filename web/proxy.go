package web

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ltto/T/tp"
)

var RouterMap = make(map[string]*RouterInfo, 0)

type RouterInfo struct {
	Title        string
	Desc         string
	Tags         []string
	Mapping      string
	HttpMethod   string
	InterfaceMap tp.H
	Do           interface{}
	f            *Func
	rout         bool
}

func (r *RouterInfo) GetMapping() string {
	return r.Mapping
}

func (r *RouterInfo) GetHttpMethod() string {
	return r.HttpMethod
}

func (r *RouterInfo) GetTitle() string {
	return r.Title
}

func (r *RouterInfo) GetDesc() string {
	return r.Desc
}

func (r *RouterInfo) GetTags() []string {
	return r.Tags
}

func (r *RouterInfo) GetInterfaceMap() map[string]interface{} {
	return r.InterfaceMap
}

func (r *RouterInfo) GetIn() *reflect.Type {
	if r.f != nil {
		return r.f.in
	}
	return nil
}

func (r *RouterInfo) GetOut() *reflect.Type {
	if r.f != nil {
		return r.f.out
	}
	return nil
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
		panic(fmt.Sprintf("重复的路由配置 `/%s__%s`", r.Mapping, r.HttpMethod))
	}
	r.f = NewFunc(r.Do)
	RouterMap[r.Mapping+r.HttpMethod] = r
	r.rout = true
	han := func(c *gin.Context) {
		if out, err := r.f.Call(NewContext(c)); err != nil {
			c.JSON(http.StatusOK, err.Error())
		} else {
			c.JSON(http.StatusOK, out)
		}
	}
	switch r.HttpMethod {
	case http.MethodGet:
		g.GET(r.Mapping, han, )
	case http.MethodPost:
		g.POST(r.Mapping, han, )
	case http.MethodPut:
		g.PUT(r.Mapping, han, )
	case http.MethodDelete:
		g.DELETE(r.Mapping, han, )
	}
	return *r
}
