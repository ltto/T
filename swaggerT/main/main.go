package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	".."
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/spec"
	"github.com/swaggo/swag"
)

func main() {

	schema := spec.Schema{
		SchemaProps: spec.SchemaProps{
			//Ref: spec.Ref{Ref: ref},
			Type: []string{"integer"},
		},
	}
	swaggerSpec := &spec.Swagger{
		//VendorExtensible: swagger.VendorExtensible,
		SwaggerProps: spec.SwaggerProps{
			Consumes: []string{"xxx", "xxxs"},
			Produces: []string{"xxx", "xxxs"},
			Swagger:  "2.0",
			Host:     "127.0.0.1",
			BasePath: "/api",
			Paths: &spec.Paths{
				Paths: map[string]spec.PathItem{},
			},
			Tags:        []spec.Tag{{TagProps: spec.TagProps{Name: "user"}}},
			Definitions: map[string]spec.Schema{"schema": schema},
		},
	}
	swaggerSpec.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Description:    "{{.Description}}",
			Title:          "{{.Title}}",
			TermsOfService: "http://我的服务器",
			Contact: &spec.ContactInfo{
				Name:  "我的邮箱",
				URL:   "http://www.tov2.com",
				Email: "121172722@qq.com",
			},
			License: &spec.License{
				Name: "我的主页",
				URL:  "http://tov2.com",
			},
			Version: "3.0",
		},
	}
	swaggerT.GetSchemaStruct(reflect.TypeOf(U{}))
	swaggerSpec.Definitions = swaggerT.SchemaMap.Map

	bytes, e := swaggerSpec.MarshalJSON()
	if e != nil {
		panic(e)
	}
	fmt.Println(ioutil.WriteFile("main6.json", bytes, 0777))
}

type U struct {
	Name string
	A    A
}
type A struct {
	Age int
}

func s(c *gin.Context) {

	type swaggerUIBundle struct {
		URL string
	}

	var matches []string
	if matches = rexp.FindStringSubmatch(c.Request.RequestURI); len(matches) != 3 {
		c.Status(404)
		c.Writer.Write([]byte("404 page not found"))
		return
	}
	path := matches[2]

	if strings.HasSuffix(path, ".html") {
		c.Header("Content-Type", "text/html; charset=utf-8")
	} else if strings.HasSuffix(path, ".css") {
		c.Header("Content-Type", "text/css; charset=utf-8")
	} else if strings.HasSuffix(path, ".js") {
		c.Header("Content-Type", "application/javascript")
	} else if strings.HasSuffix(path, ".json") {
		c.Header("Content-Type", "application/json")
	}

	switch path {
	case "index.html":
		index.Execute(c.Writer, &swaggerUIBundle{
			URL: config.URL,
		})
	case "doc.json":
		doc, err := swag.ReadDoc()
		if err != nil {
			panic(err)
		}
		c.Writer.Write([]byte(doc))
		return
	default:
		h.ServeHTTP(c.Writer, c.Request)
	}
}
