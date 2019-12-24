package webT

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)


func GetSession(c gin.Context) *sessions.Session {
	//get, err := session.Get("sessionId", c)
	//if err != nil {
	//	panic(err)
	//}
	//return get
	return nil
}
