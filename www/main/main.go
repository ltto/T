package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ltto/T/www"
)

func main() {
	www.Engine.Use(func(c *gin.Context) {
		fmt.Println("前置")
		c.Next()
		_, exists := c.Get("su")
		if exists {
			fmt.Println("后置")
		}
	})
	www.GetMapping("/hello", func(c *www.Context) interface{} {
		fmt.Println("ok")
		c.Context.Set("su", true)
		return gin.H{"hello": true}
	})

	www.Engine.Run(":8080")
}
