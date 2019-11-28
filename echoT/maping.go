package echoT

import (
	"github.com/labstack/echo/v4"
)

var e = echo.New()
var g = e.Group("/")

func init() {
	e.Binder = &MyBinder{}
	AddMiddleware(sessionMiddleware)
	e.HTTPErrorHandler = ErrorHandler
}

func AddMiddleware(m ...echo.MiddlewareFunc) {
	g = e.Group("/", m...)
}

func Run(address string) error {
	server := swagger()
	g.GET("swagger/*", func(c echo.Context) error {
		return server.Http(c.Request(), c.Response())
	})
	return e.Start(address)
}
