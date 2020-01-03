package web

import (
	"github.com/gin-gonic/gin"
	"github.com/ltto/T/swaggerT"
)

var e = gin.New()
var g = e.Group("/")

func init() {
}

func Run(address string) error {
	server := swaggerT.Server{
		WWW:      "http://127.0.0.1",
		Title:    "项目-swagger",
		Desc:     "项目描述",
		BasePath: "/",
		Host:     "127.0.0.1:8080",
		Version:  "1.0",
	}
	var routerList = make([]swaggerT.Router, 0, len(RouterMap))
	for k := range RouterMap {
		routerList = append(routerList, RouterMap[k])
	}
	server.SwaggerList(routerList...)
	g.GET("swagger/:path", func(c *gin.Context) {
		server.Http(c.Request, c.Writer)
	})
	return e.Run(address)
}
