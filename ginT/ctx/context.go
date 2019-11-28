package ctx

import (
	"encoding/json"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ltto/T/ginT/vo"
)

type Context struct {
	*gin.Context
}

func (c *Context) Session() sessions.Session {
	return sessions.Default(c.Context)
}

func (c *Context) JsonBody(data interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(data)
}

func (c *Context) GetPage() *vo.Page {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	return vo.NewPage(page, size)
}

func (c *Context) SetHtmlTemp(data interface{}) {
	c.Set("temp", data)
}

func (c *Context) AddHtmlParam(key string, data interface{}) {
	if _, exists := c.Get("html"); !exists {
		c.Set("html", make(map[string]interface{}))
	}
	get, _ := c.Get("html")
	m := get.(map[string]interface{})
	m[key] = data
}
