package echoT

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var sessionMiddleware = session.Middleware(sessions.NewCookieStore([]byte("133")))

func GetSession(c echo.Context) (*sessions.Session) {
	get, err := session.Get("sessionId", c)
	if err != nil {
		panic(err)
	}
	return get
}
