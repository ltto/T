package webT

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/ltto/T/gobox/ref"
)

type MyBinder struct {
	echo.DefaultBinder
}

func (b *MyBinder) Bind(i interface{}, c echo.Context) (err error) {
	paths := make(map[string][]string)
	for _, v := range c.ParamNames() {
		paths[v] = []string{c.Param(v)}
	}
	if err = ref.BindDataStr(i, paths, "path", false); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	req := c.Request()
	if req.ContentLength == 0 {
		if req.Method == http.MethodGet || req.Method == http.MethodDelete {
			if err = ref.BindDataStr(i, c.QueryParams(), "query", true); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
			}
			return
		}
		return echo.NewHTTPError(http.StatusBadRequest, "Request body can't be empty")
	}
	ctype := req.Header.Get(echo.HeaderContentType)
	switch {
	case strings.HasPrefix(ctype, echo.MIMEApplicationJSON):
		if err = json.NewDecoder(req.Body).Decode(i); err != nil {
			if ute, ok := err.(*json.UnmarshalTypeError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
			} else if se, ok := err.(*json.SyntaxError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
			}
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
	case strings.HasPrefix(ctype, echo.MIMEApplicationXML), strings.HasPrefix(ctype, echo.MIMETextXML):
		if err = xml.NewDecoder(req.Body).Decode(i); err != nil {
			if ute, ok := err.(*xml.UnsupportedTypeError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unsupported type error: type=%v, error=%v", ute.Type, ute.Error())).SetInternal(err)
			} else if se, ok := err.(*xml.SyntaxError); ok {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: line=%v, error=%v", se.Line, se.Error())).SetInternal(err)
			}
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
	case strings.HasPrefix(ctype, echo.MIMEApplicationForm), strings.HasPrefix(ctype, echo.MIMEMultipartForm):
		params, err := c.FormParams()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
		if err = ref.BindDataStr(i, params, "form", true); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
		}
	default:
		return echo.ErrUnsupportedMediaType
	}
	return
}
