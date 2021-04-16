package node

type CodeToken struct {
	tag   string
	data  string
	child []Node
	attr  map[string]string
}

func (c *CodeToken) Tag() string {
	return c.tag
}

func (c *CodeToken) Data() string {
	return c.data
}

func (c *CodeToken) Child() []Node {
	return c.child
}

func (c *CodeToken) Attr(key string) string {
	return c.attr[key]
}
func Select(c ...*CodeToken) *CodeToken {
	token := &CodeToken{tag: "select"}
	for i := range c {
		token.child = append(token.child, doPareChild(c[i])...)
	}
	return token
}
func IF_(test string, c ...*CodeToken) *CodeToken {
	token := &CodeToken{tag: "if", attr: map[string]string{"test": test}}
	for i := range c {
		token.child = append(token.child, doPareChild(c[i])...)
	}
	return token
}

type F struct {
	Index      string
	Item       string
	Collection string
	Open       string
	Separator  string
	Close      string
}

func For_(f F, c ...*CodeToken) *CodeToken {
	token := &CodeToken{tag: "foreach", attr: map[string]string{
		"index":      f.Index,
		"item":       f.Item,
		"collection": f.Collection,
		"open":       f.Open,
		"separator":  f.Separator,
		"close":      f.Close,
	}}
	for i := range c {
		token.child = append(token.child, doPareChild(c[i])...)
	}
	return token
}
func Include_(refId string, c ...*CodeToken) *CodeToken {
	token := &CodeToken{tag: "sql", attr: map[string]string{"refid": refId}}
	for i := range c {
		token.child = append(token.child, doPareChild(c[i])...)
	}
	return token
}
func Text_(data string) *CodeToken {
	return &CodeToken{data: data}
}
