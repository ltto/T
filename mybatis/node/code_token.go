package node

type CodeToken struct {
	string
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
func IncludeSQL(id string, c ...Token) Token {
	token := &CodeToken{tag: "sql", attr: map[string]string{"id": id}}
	for i := range c {
		token.child = append(token.child, doPareChild(c[i])...)
	}
	return token
}
func Select(c ...Token) Token {
	token := &CodeToken{tag: "select"}
	for i := range c {
		token.child = append(token.child, doPareChild(c[i])...)
	}
	return token
}
func Insert(c ...Token) Token {
	token := &CodeToken{tag: "insert"}
	for i := range c {
		token.child = append(token.child, doPareChild(c[i])...)
	}
	return token
}
func Update(c ...Token) Token {
	token := &CodeToken{tag: "update"}
	for i := range c {
		token.child = append(token.child, doPareChild(c[i])...)
	}
	return token
}
func Delete(c ...Token) Token {
	token := &CodeToken{tag: "delete"}
	for i := range c {
		token.child = append(token.child, doPareChild(c[i])...)
	}
	return token
}
func IF_(test string, c ...Token) Token {
	token := &CodeToken{tag: "if", attr: map[string]string{"test": test}}
	for i := range c {
		token.child = append(token.child, doPareChild(c[i])...)
	}
	return token
}

func For_(index string, item string, collection string, open string, separator string, close string, c ...Token) Token {
	token := &CodeToken{tag: "foreach", attr: map[string]string{
		"index":      index,
		"item":       item,
		"collection": collection,
		"open":       open,
		"separator":  separator,
		"close":      close,
	}}
	for i := range c {
		token.child = append(token.child, doPareChild(c[i])...)
	}
	return token
}
func Include_(refId string) Token {
	return &CodeToken{tag: "include", attr: map[string]string{"refid": refId}}
}
func Text_(data string) Token {
	return &CodeToken{data: data}
}
