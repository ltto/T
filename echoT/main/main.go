package main

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ltto/T/echoT"
	"github.com/ltto/T/echoT/vo"
)

func main() {
	echoT.R(echoT.RouterInfo{Mapping: "/user/:id", HttpMethod: http.MethodGet,
		Auth:         false,
		InterfaceMap: echoT.InterfaceMap{"data": reflect.TypeOf(User{})},
		Do: func(res struct {
			Name123 string `query:"name123"`
			ID      string `path:"id"`
			User
		}, c echo.Context) vo.Result {
			sess := echoT.GetSession(c)
			defer func() { sess.Save(c.Request(), c.Response()) }()
			fmt.Println(sess.Values["user"], sess.IsNew)
			unix := time.Now().Unix()
			sess.Values["user"] = unix
			fmt.Println(res)
			return vo.Result{}
		},
	})

	fmt.Println(echoT.Run(":8080"))
}

type User struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}
