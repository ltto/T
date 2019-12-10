package param

import "reflect"

type ContextType string

type Router struct {
	method string
	path   string
	req    ContextType
	params Params
	reps   ContextType
	view   reflect.Type
	Func   interface{}
}
