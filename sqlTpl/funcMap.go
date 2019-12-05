package sqlTpl

import (
	"fmt"
	"reflect"

	"github.com/guregu/null"
)

var funcMap = map[string]interface{}{
	"tplfor":  Foreach,
	"blank":   Blank,
	"unBlank": unBlank,
}

func Foreach(data interface{}, open string, close string, separator string, val string) string {
	v := reflect.ValueOf(data)
	for v.Kind() == reflect.Ptr {
		v.Elem()
	}
	strs := open
	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Cap(); i++ {
			strs += fmt.Sprintf("#{%s[%d]}", val, i)
			if i != v.Cap()-1 {
				strs += separator
			}
		}
	case reflect.Map:
		keys := v.MapKeys()
		for _, k := range keys {
			strs += fmt.Sprintf("#{%s.%s}", val, k.String()) + separator
		}
		strs = strs[0 : len(strs)-len(separator)]
	default:
		panic("need [] ro map")
	}
	return strs + close
}

func Blank(data interface{}) bool {
	if data == nil {
		return true
	}
	switch d := data.(type) {
	case string:
		return d == ""
	case null.Int:
		return !d.Valid
	case null.Bool:
		return !d.Valid
	case null.Float:
		return !d.Valid
	case null.String:
		return !d.Valid
	case null.Time:
		return !d.Valid
	}
	return false
}

func unBlank(data interface{}) bool {
	return !Blank(data)
}
