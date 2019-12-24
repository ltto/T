package webT

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type Context struct {
	MyBinder
	*gin.Context
	ses *sessions.Session
}

func (c *Context) Bind(ptr interface{}) error{
	return c.MyBinder.Bind(ptr, c)
}
func NewContext(context *gin.Context) *Context {
	return &Context{Context: context}
}

func (c Context) Session() *sessions.Session {
	return c.ses
}
