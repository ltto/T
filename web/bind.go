package web

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ltto/T/gobox/ref"
)

type MyBinder struct {
}

func (b *MyBinder) Bind(i interface{}, c *Context) (err error) {
	paths := make(map[string][]string)
	var paramNames []string
	for _, param := range c.Params {
		paramNames = append(paramNames, param.Key)
	}
	for _, v := range paramNames {
		paths[v] = []string{c.Param(v)}
	}
	if err = ref.BindDataStr(i, paths, "path", false); err != nil {
		return NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := c.Request
	if req.ContentLength == 0 {
		if req.Method == http.MethodGet || req.Method == http.MethodDelete {
			if err = ref.BindDataStr(i, c.Request.URL.Query(), "query", true); err != nil {
				return NewHTTPError(http.StatusBadRequest, err.Error())
			}
			return
		}

		return NewHTTPError(http.StatusBadRequest, "Request body can't be empty")
	}
	ctype := req.Header.Get("Content-Type")
	switch {
	case strings.HasPrefix(ctype, gin.MIMEJSON):
		if err = json.NewDecoder(req.Body).Decode(i); err != nil {
			if ute, ok := err.(*json.UnmarshalTypeError); ok {
				return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset))
			} else if se, ok := err.(*json.SyntaxError); ok {
				return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error()))
			}
			return NewHTTPError(http.StatusBadRequest, err.Error())
		}
	case strings.HasPrefix(ctype, gin.MIMEXML), strings.HasPrefix(ctype, gin.MIMEXML2):
		if err = xml.NewDecoder(req.Body).Decode(i); err != nil {
			if ute, ok := err.(*xml.UnsupportedTypeError); ok {
				return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unsupported type error: type=%v, error=%v", ute.Type, ute.Error()))
			} else if se, ok := err.(*xml.SyntaxError); ok {
				return NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: line=%v, error=%v", se.Line, se.Error()))
			}
			return NewHTTPError(http.StatusBadRequest, err.Error())
		}
	case strings.HasPrefix(ctype, gin.MIMEPOSTForm), strings.HasPrefix(ctype, gin.MIMEMultipartPOSTForm):
		params := c.Request.PostForm
		if err = ref.BindDataStr(i, params, "form", true); err != nil {
			return NewHTTPError(http.StatusBadRequest, err.Error())
		}
	default:
		return errors.New("err unsupported media type")
	}
	return
}
