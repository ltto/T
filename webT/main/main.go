package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ltto/T/webT"
	"github.com/ltto/T/webT/vo"
)

func main() {

	webT.R(webT.RouterInfo{Mapping: "/user/:id", HttpMethod: http.MethodGet,
		Auth: false,
		Doc: webT.RouterDoc{
			Desc:  "简介",
			Title: "获取用户",
			Tags:  []string{"user"},
		},
		InterfaceMap: webT.InterfaceMap{"data": User{}},
		Do: func(res struct {
			Name123 string `query:"name123"`
			ID      string `path:"id"`
			//User
		}, c *gin.Context) vo.Result {
			fmt.Println(res)
			return vo.Result{}
		},
	})

	fmt.Println(webT.Run(":8080"))
}

type User struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}
