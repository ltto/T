package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ltto/T/swaggerT"
)

var e = gin.New()
var g = e.Group("/")

func init() {
}

type Docer interface {
	Http(r *http.Request, w http.ResponseWriter) error
	Init(map[string]*RouterInfo)
}

func Run(address string, server *swaggerT.Server) error {
	if server != nil {
		var routerList = make([]swaggerT.Router, 0, len(RouterMap))
		for k := range RouterMap {
			routerList = append(routerList, RouterMap[k])
		}
		server.SwaggerList(routerList...)
		g.GET("swagger/:path", func(c *gin.Context) {
			server.Http(c.Request, c.Writer)
		})
	}

	return e.Run(address)
}
