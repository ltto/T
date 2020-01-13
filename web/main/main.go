package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ltto/T/swaggerT"
	"github.com/ltto/T/tp"
	"github.com/ltto/T/web"
	"github.com/ltto/T/web/vo"
)

func main() {

	web.R(web.RouterInfo{Mapping: "/user/:id", HttpMethod: http.MethodGet,
		Desc:         "简介",
		Title:        "获取用户",
		Tags:         []string{"user"},
		InterfaceMap: tp.H{"data": User{}},
		Do: func(res struct {
			Name123 string `query:"name123"`
			ID      string `path:"id"`
			User
		}, c *gin.Context) vo.Result {
			fmt.Println(res)
			return vo.Result{}
		},
	})

	server := swaggerT.Server{
		WWW:      "http://127.0.0.1",
		Title:    "项目-swagger",
		Desc:     "项目描述",
		BasePath: "/",
		Host:     "127.0.0.1:8080",
		Version:  "1.0",
	}
	fmt.Println(web.Run(":8080", &server))
}

type User struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}
