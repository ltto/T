package swaggerT

import (
	"net/http"
	"reflect"

	"github.com/go-openapi/spec"
)

type ParamType string
type IN string
type TPaths struct {
	spec.Paths
}

const (
	ParaArray   ParamType = "array"
	ParaBoolean ParamType = "boolean"
	ParaInteger ParamType = "integer"
	ParaNumber  ParamType = "number"
	ParaObject  ParamType = "object"
	ParaString  ParamType = "string"

	InBody     IN = "body"
	InHeader   IN = "header"
	InFormData IN = "formData"
	InQuery    IN = "query"
	InPath     IN = "path"
)

func NewTPaths() *TPaths {
	return &TPaths{
		Paths: spec.Paths{
			Paths: map[string]spec.PathItem{},
		},
	}
}

func (T *TPaths) AppendPath(path Path) *spec.Operation {
	if _, ok := T.Paths.Paths[path.URL]; !ok {
		T.Paths.Paths[path.URL] = spec.PathItem{}
	}
	item := T.Paths.Paths[path.URL]
	operation := spec.Operation{
		OperationProps: spec.OperationProps{
			Description: path.Desc,
			Consumes:    []string{},
			Produces:    []string{},
			Schemes:     []string{},
			Tags:        path.Tags,
			Summary:     path.Title,
			ID:          "",
			Deprecated:  false, //不推荐使用
			Security:    nil,
			Parameters:  []spec.Parameter{},
			Responses:   &spec.Responses{},
		},
	}
	path.InitParam(&operation)
	operation.Responses.Default = path.GetResponse()
	switch path.Method {
	case http.MethodGet:
		item.Get = &operation
	case http.MethodPost:
		item.Post = &operation
	case http.MethodPut:
		item.Put = &operation
	case http.MethodDelete:
		item.Delete = &operation
	case http.MethodOptions:
		item.Options = &operation
	case http.MethodHead:
		item.Head = &operation
	case http.MethodPatch:
		item.Patch = &operation
	}
	T.Paths.Paths[path.URL] = item
	return &operation
}

func (p Path) InitParam(operation *spec.Operation) {
	for name, val := range p.Params.Form {
		parameter := p.initNoBody(name, val, InFormData)
		operation.Parameters = append(operation.Parameters, parameter)
	}
	for name, val := range p.Params.Query {
		parameter := p.initNoBody(name, val, InQuery)
		operation.Parameters = append(operation.Parameters, parameter)
	}
	for name, val := range p.Params.Path {
		parameter := p.initNoBody(name, val, InPath)
		operation.Parameters = append(operation.Parameters, parameter)
	}
	if len(p.Params.Body) > 0 {
		parameter := p.initBody(p.Params.Body)
		operation.Parameters = append(operation.Parameters, parameter)
	}
}

func (p Path) initNoBody(name string, val reflect.Type, in IN) spec.Parameter {
	paramType, Format := getParamType(val)
	parameter := spec.Parameter{}
	parameter.Name = name
	parameter.In = string(in)
	ref := SchemaMap.getRef(NewKey(val, InterfaceMap{}))
	parameter.Schema = ref
	parameter.Type = string(paramType)
	parameter.Format = Format
	return parameter
}

func (p Path) initBody(m map[string]reflect.Type) spec.Parameter {
	parameter := spec.Parameter{}
	parameter.Name = "body"
	parameter.In = string(InBody)
	schema := spec.Schema{}
	schema.Properties = make(map[string]spec.Schema)
	for name, val := range m {
		ref := SchemaMap.getRef(NewKey(val, InterfaceMap{}))
		if ref != nil {
			schema.Properties[name] = *ref
		}
	}
	key := NewKey(reflect.TypeOf(struct{}{}), InterfaceMap{})
	key.K = p.URL + "_" + p.Method
	SchemaMap.Set(key, &schema)
	ref := SchemaMap.getRef(key)
	if ref == nil {
		panic("ref is nil")
	}
	parameter.Schema = ref
	parameter.Type = string(ParaObject)
	parameter.Format = string(ParaObject)
	return parameter
}

func (p Path) GetResponse() *spec.Response {
	props := spec.Response{}
	props.Description = "Response"
	if p.Out != nil {
		ref := SchemaMap.getRef(NewKey(*p.Out, p.InterfaceMap))
		schema := ref
		props.Schema = schema
	}
	return &props
}
