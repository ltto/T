package www

import "net/http"

type ErrHandler func(c *Context, err error)

var DefaultErrHandler = func(c *Context, err error) {
	c.String(http.StatusInternalServerError, err.Error())
	c.Abort()
}
