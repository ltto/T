package main

import (
	"github.com/ltto/T/www"
	"github.com/ltto/T/www/rest"
)

func main() {
	www.Get("/hello", func(c *www.Context) interface{} {
		return rest.ReturnFile("/Users/liutongtong/gocode/T/www/main/rgb.png")
	})

	www.Run()
}
