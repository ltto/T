package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ltto/T/www"
	"github.com/ltto/T/www/rest"
)

func main() {
	www.Get("/hello", func(c *www.Context) interface{} {
		return rest.ReturnRedirect("https://www.baidu.com")
	})

	www.Get("/ok", func(c *www.Context) string {
		return "OK"
	})
	www.Get("/json", func(c *www.Context) gin.H {
		return gin.H{"json": "ok"}
	})
	www.Run()
}
