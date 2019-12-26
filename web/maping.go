package web

import (
	"github.com/gin-gonic/gin"
)

var e = gin.New()
var g = e.Group("/")

func init() {
}

func Run(address string) error {
	server := swagger("/")
	g.GET("swagger/:path", func(c *gin.Context) {
		server.Http(c.Request, c.Writer)
	})
	return e.Run(address)
}
