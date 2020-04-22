package main

import (
	"github.com/ltto/T/www"
)

func main() {
	www.GetMapping("/hello", func(c *www.Context) interface{} {
		return www.ReturnFile("www/main/rgb.png")
	})

	www.Engine.Run(":8080")
}
