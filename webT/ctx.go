package webT

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	ses *sessions.Session
}

func (c Context) Session() *sessions.Session {
	return c.ses
}
