package echoT

import (
	"reflect"

	"github.com/ltto/T/swaggerT"
)

type Params struct {
	Form  map[string]reflect.Type
	Path  map[string]reflect.Type
	Query map[string]reflect.Type
	Body  map[string]reflect.Type
}

func swagger() swaggerT.Server {
	var paths []swaggerT.Path
	for _, v := range RouterMap {
		params := swaggerT.Params{}
		if v.f.in != nil {
			params.Scanner(*v.f.in)
		}
		paths = append(paths, swaggerT.Path{
			Tags:         []string{"user"},
			Desc:         "接口描述",
			Title:        "接口描述",
			URL:          "/" + v.Mapping,
			Method:       v.HttpMethod,
			Out:          v.f.out,
			Params:       params,
			InterfaceMap: swaggerT.InterfaceMap(v.InterfaceMap),
		})
	}
	server := swaggerT.Server{
		Tags:     []string{"user"},
		WWW:      "http://127.0.0.1",
		Title:    "title",
		Desc:     "项目描述",
		BasePath: "/",
		Host:     "127.0.0.1:8080",
		Version:  "1.0",
		Paths:    paths,
	}
	server.Init()
	return server
}
