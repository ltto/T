package web

import (
	"reflect"
)

type Params struct {
	Form  map[string]reflect.Type
	Path  map[string]reflect.Type
	Query map[string]reflect.Type
	Body  map[string]reflect.Type
}

