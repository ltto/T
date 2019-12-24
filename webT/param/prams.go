package param

type Params struct {
	body     map[string]Param
	path     map[string]Param
	formData map[string]Param
	query    map[string]Param
	head     map[string]Param
	objs     map[string]interface{}
}

func (p *Params) Append(in, key string, pa Param) {
	if p.body == nil {
		p.body = make(map[string]Param)
	}
	if p.path == nil {
		p.path = make(map[string]Param)
	}
	if p.formData == nil {
		p.formData = make(map[string]Param)
	}
	if p.query == nil {
		p.query = make(map[string]Param)
	}
	if p.head == nil {
		p.head = make(map[string]Param)
	}
	switch in {
	case "body":
		p.body[key] = pa
	case "formData":
		p.formData[key] = pa
	case "path":
		p.path[key] = pa
	case "query":
		p.query[key] = pa
	case "head":
		p.head[key] = pa
	}
}
