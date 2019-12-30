package www

import "net/http"

type ErrHandler interface {
	Handler(c Context, err error)
}

var DefaultErrHandler = &DefErrHandler{}

type DefErrHandler struct {
}

func (d *DefErrHandler) Handler(c *Context, err error) {
	c.String(http.StatusInternalServerError, err.Error())
	c.Abort()
}
