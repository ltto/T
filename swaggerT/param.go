package swaggerT

import (
	"reflect"
	"strings"

	"github.com/go-openapi/jsonreference"
	"github.com/go-openapi/spec"
	"github.com/ltto/T/gobox/ref"
)

type schema struct {
	Map    map[string]spec.Schema
	RefMap map[string]*spec.Schema
}
type Key struct {
	K string
	t reflect.Type
	m InterfaceMap
}

func NewKey(t reflect.Type, m InterfaceMap) Key {
	return Key{t: t, m: m}
}

func (k Key) FullName() string {
	if k.K != "" {
		return strings.ReplaceAll(k.K,"/","_")
	}
	Generic := ""
	if len(k.m) > 0 {
		for e := range k.m {
			Generic += strings.ReplaceAll(ref.FullName(k.m[e]), "/", "_") + ","
		}
	} else {
		Generic = ","
	}
	return strings.ReplaceAll(ref.FullName(k.t), "/", "_") + "<" + Generic[:len(Generic)-1] + ">"
}

func (s *schema) getRef(k Key) *spec.Schema {
	key := k.FullName()
	_, ok := s.RefMap[key]
	if !ok && k.t != nil {
		SchemaMap.Set(k, byT(k, ""))
	}
	return s.RefMap[key]
}
func (s *schema) Set(k Key, v *spec.Schema) {
	if v == nil {
		return
	}
	key := k.FullName()
	if !ref.IsBaseTime(k.t) {
		s.Map[key] = spec.Schema{SchemaProps: spec.SchemaProps{
			Type:       v.Type,
			Properties: v.Properties,
		}}
		refObj, e := jsonreference.New("#/definitions/" + key)
		if e != nil {
			panic(e)
		}
		s.RefMap[key] = &spec.Schema{SchemaProps: spec.SchemaProps{
			Type: v.Type,
			Ref:  spec.Ref{Ref: refObj},
		}}
	} else {
		s.RefMap[key] = v
	}
}

var SchemaMap = schema{
	Map:    make(map[string]spec.Schema, 0),
	RefMap: make(map[string]*spec.Schema, 0),
}
var SchemaObj = spec.Schema{SchemaProps: spec.SchemaProps{
	Type:   []string{string(ParaObject)},
	Format: "interface{}",
}}
var SchemaTime = spec.Schema{SchemaProps: spec.SchemaProps{
	Type:   []string{string(ParaObject)},
	Format: "Time",
}}

func byT(k Key, field string) *spec.Schema {
	t := k.t
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.String() {
	case "null.Bool":
		return getSchemaBse(reflect.TypeOf(true))
	case "null.Int":
		return getSchemaBse(reflect.TypeOf(int64(0)))
	case "null.Float":
		return getSchemaBse(reflect.TypeOf(float64(0.1)))
	case "null.Time", "time.Time":
		return &SchemaTime
	case "null.String":
		return getSchemaBse(reflect.TypeOf("string"))
	default:
		switch t.Kind() {
		case reflect.Bool:
			return getSchemaBse(t)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return getSchemaBse(t)
		case reflect.Float32, reflect.Float64:
			return getSchemaBse(t)
		case reflect.String:
			return getSchemaBse(t)
		case reflect.Slice:
			return getSchemaArr(NewKey(t, k.m))
		case reflect.Struct:
			return getSchemaStruct(NewKey(t, k.m))
		case reflect.Interface:
			if k.m != nil {
				if i, ok := k.m[field]; ok {
					if i.Kind() == reflect.Ptr {
						i = i.Elem()
					}
					switch i.Kind() {
					case reflect.Interface, reflect.Struct:
						return getSchemaStruct(NewKey(i, k.m))
					case reflect.Slice:
						return getSchemaArr(NewKey(i, k.m))
					}
				}
			}
		}
		return &SchemaObj
	}
}
func getSchemaStruct(k Key) *spec.Schema {
	schema := spec.Schema{}
	t := k.t
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("滚蛋")
	}

	schema.Properties = make(map[string]spec.Schema)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		ft := field.Type
		if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
		}
		name := field.Tag.Get("json")
		if name == "" {
			name = field.Name
		}
		b := byT(NewKey(ft, k.m), name)
		if b != nil {
			schema.Properties[name] = *b
		}
	}

	return &schema
}

func getSchemaBse(t reflect.Type) *spec.Schema {
	schema := spec.Schema{SchemaProps: spec.SchemaProps{}}
	paramType, Format := getParamType(t)
	schema.Type = []string{string(paramType)}
	schema.Format = Format
	return &schema
}
func getSchemaArr(k Key) *spec.Schema {
	t := k.t
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	schema := spec.Schema{}
	schema.Type = []string{string(ParaArray)}
	array := spec.SchemaOrArray{Schema: byT(NewKey(t.Elem(), k.m), "")}
	schema.Items = &array
	return &schema
}

func getParamType(t reflect.Type) (paramType ParamType, Format string) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.String() {
	case "null.Bool":
		return ParaBoolean, t.Name()
	case "null.Int":
		return ParaInteger, t.Name()
	case "null.Float":
		return ParaNumber, t.Name()
	case "null.Time", "time.Time":
		return ParaObject, t.Name()
	case "null.String":
		return ParaString, t.Name()
	default:
		switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return ParaInteger, t.Name()
		case reflect.String:
			return ParaString, t.Name()
		case reflect.Float32, reflect.Float64:
			return ParaNumber, t.Name()
		case reflect.Bool:
			return ParaBoolean, t.Name()
		case reflect.Slice:
			return ParaArray, t.Name()
		default:
			return ParaObject, t.Name()
		}
	}
}
