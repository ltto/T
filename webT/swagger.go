package webT

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

func swagger(baseURL string, ) swaggerT.Server {
	server := swaggerT.Server{
		WWW:      "http://127.0.0.1",
		Title:    "项目-swagger",
		Desc:     "项目描述",
		BasePath: baseURL,
		Host:     "127.0.0.1:8080",
		Version:  "1.0",
	}
	var paths []swaggerT.Path
	var AllTags = make(map[string]struct{})
	for _, v := range RouterMap {
		params := swaggerT.Params{}
		if v.f.in != nil {
			params.Scanner(*v.f.in)
		}
		if len(v.Doc.Tags) == 0 {
			v.Doc.Tags = []string{"base"}
		}
		for i := range v.Doc.Tags {
			AllTags[v.Doc.Tags[i]] = struct{}{}
		}
		TMap := make(map[string]reflect.Type)
		for k, v := range v.InterfaceMap {
			if v == nil {
				panic("InterfaceMap val is nil")
			}
			vt := reflect.TypeOf(v)
			for vt.Kind() == reflect.Ptr {
				vt = vt.Elem()
			}
			TMap[k] = vt
		}
		paths = append(paths, swaggerT.Path{
			Tags:         v.Doc.Tags,
			Desc:         v.Doc.Desc,
			Title:        v.Doc.Title,
			URL:          server.BasePath + v.Mapping,
			Method:       v.HttpMethod,
			Out:          v.f.out,
			Params:       params,
			InterfaceMap: TMap,
		})
	}
	tagsArr := make([]string, 0, len(AllTags))
	for k := range AllTags {
		tagsArr = append(tagsArr, k)
	}
	server.Paths = paths
	server.Tags = tagsArr
	server.Init()
	return server
}
