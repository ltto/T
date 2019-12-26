package main

import (
	"fmt"

	"github.com/ltto/T/gobox/str"
)

func main() {
	expand := str.ExpandS("{s}12345678910{sss}", func(s string) string {
		fmt.Println(s)
		return "ok"
	})
	fmt.Println(expand)
}
