package swaggerT

import (
	"github.com/go-openapi/spec"
)

func GetInfo(ser Server) []byte {
	schema := spec.Schema{
		SchemaProps: spec.SchemaProps{
			//Ref: spec.Ref{Ref: ref},
			Type: []string{"integer"},
		},
	}
	var tags []spec.Tag
	for _, v := range ser.Tags {
		tags = append(tags, spec.Tag{TagProps: spec.TagProps{Name: v}})
	}
	swaggerSpec := &spec.Swagger{
		//VendorExtensible: swagger.VendorExtensible,
		SwaggerProps: spec.SwaggerProps{
			Consumes: []string{"http", "https"},
			Produces: []string{"http", "https"},
			Swagger:  "2.0",
			Host:     ser.Host,
			BasePath: ser.BasePath,
			Paths: &spec.Paths{
				Paths: map[string]spec.PathItem{},
			},
			Tags:        tags,
			Definitions: map[string]spec.Schema{"schema": schema},
		},
	}
	swaggerSpec.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Description:    ser.Desc,
			Title:          ser.Title,
			TermsOfService: ser.WWW,
			Contact: &spec.ContactInfo{
				Name:  "我的邮箱",
				URL:   "http://www.tov2.com",
				Email: "121172722@qq.com",
			},
			License: &spec.License{
				Name: "我的主页",
				URL:  "http://tov2.com",
			},
			Version: ser.Version,
		},
	}
	paths := NewTPaths()
	for _, path := range ser.Paths {
		paths.AppendPath(path)
	}
	swaggerSpec.Paths.Paths = paths.Paths.Paths
	swaggerSpec.Definitions = SchemaMap.Map

	bytes, e := swaggerSpec.MarshalJSON()
	if e != nil {
		panic(e)
	}
	return bytes
}
