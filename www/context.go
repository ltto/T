package www

import (
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	sessions.Session
}

func (c *Context) JsonBody(data interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(data)
}

func (c *Context) SetHtmlTemp(data interface{}) {
	c.Context.Set("temp", data)
}

func (c *Context) AddCParam(key string, data interface{}) {
	if _, exists := c.Context.Get("_param"); !exists {
		c.Context.Set("_param", make(map[string]interface{}))
	}
	get, _ := c.Context.Get("_param")
	m := get.(map[string]interface{})
	m[key] = data
}

func (c *Context) CParams() map[string]interface{} {
	get, ok := c.Context.Get("_param")
	if !ok {
		return make(map[string]interface{})
	}
	if m, ok := get.(map[string]interface{}); ok {
		return m
	} else {
		return make(map[string]interface{})
	}
}

func (c *Context) GetCParam(key string) (interface{}, bool) {
	if _, exists := c.Context.Get("_param"); !exists {
		return nil, false
	}
	get, ok := c.Context.Get("_param")
	if !ok {
		return nil, false
	}
	if m, ok := get.(map[string]interface{}); ok {
		i, exist := m[key]
		return i, exist
	} else {
		return nil, false
	}
}
