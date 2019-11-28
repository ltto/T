package ginT

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ltto/T/ginT/ctx"

	"github.com/gin-contrib/sessions"
)

const (
	sessionKeyLastAction = "lastAction"
	sessionKeyActioned   = "actioned"
	SessionID            = "SESSIONIDS"
)

var (
	Expired                        = time.Minute * 30
	SessionInterceptor Interceptor = func(c *ctx.Context) *InterceptErr {
		session := sessions.Default(c.Context)
		session.Set(sessionKeyActioned, true)
		preUnix, ok := session.Get(sessionKeyLastAction).(int64)
		if !ok {
			preUnix = int64(0)
		}
		unix := time.Now().Unix()
		session.Set(sessionKeyLastAction, unix)
		if unix-preUnix > int64(Expired.Seconds()) {
			//Expired
			if err := session.Save(); err != nil {
				return NewInterceptErr(http.StatusInternalServerError, err)
			}
		}
		s, _ := c.Cookie(SessionID)
		if s == "" {
			if err := session.Save(); err != nil {
				c.Done()
				return NewInterceptErr(http.StatusInternalServerError, err)
			}
		}
		fmt.Println(c.Cookie(SessionID))
		return NewInterceptOK()
	}
)
