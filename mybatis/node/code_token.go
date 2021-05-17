package node

import (
	"errors"
	"fmt"
)

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
func IncludeSQL(id string, c ...interface{}) Token {
	token := &CodeToken{tag: "sql", attr: map[string]string{"id": id}}

	for _, t := range toToken(c...) {
		token.child = append(token.child, doPareChild(t)...)
	}
	return token
}
func Select(includes map[string]Token, c ...interface{}) *DMLRoot {
	token := &CodeToken{tag: "select"}
	for _, t := range toToken(c...) {
		token.child = append(token.child, doPareChild(t)...)
	}
	return NewNodeRoot(token, includes)
}
func Insert(includes map[string]Token, c ...interface{}) *DMLRoot {
	token := &CodeToken{tag: "insert"}
	for _, t := range toToken(c...) {
		token.child = append(token.child, doPareChild(t)...)
	}
	return NewNodeRoot(token, includes)
}
func Update(includes map[string]Token, c ...interface{}) *DMLRoot {
	token := &CodeToken{tag: "update"}
	for _, t := range toToken(c...) {
		token.child = append(token.child, doPareChild(t)...)
	}
	return NewNodeRoot(token, includes)
}
func Delete(includes map[string]Token, c ...interface{}) *DMLRoot {
	token := &CodeToken{tag: "delete"}
	for _, t := range toToken(c...) {
		token.child = append(token.child, doPareChild(t)...)
	}
	return NewNodeRoot(token, includes)
}
func IF_(test string, c ...interface{}) Token {
	token := &CodeToken{tag: "if", attr: map[string]string{"test": test}}
	for _, t := range toToken(c...) {
		token.child = append(token.child, doPareChild(t)...)
	}
	return token
}

func For_(index string, item string, collection string, open string, separator string, close string, c ...interface{}) Token {
	token := &CodeToken{tag: "foreach", attr: map[string]string{
		"index":      index,
		"item":       item,
		"collection": collection,
		"open":       open,
		"separator":  separator,
		"close":      close,
	}}
	for _, t := range toToken(c...) {
		token.child = append(token.child, doPareChild(t)...)
	}
	return token
}
func Include_(refId string) Token {
	return &CodeToken{tag: "include", attr: map[string]string{"refid": refId}}
}
func Text_(data string) Token {
	return &CodeToken{data: data}
}
func toToken(params ...interface{}) (ts []Token) {
	for i := range params {
		switch p := params[i].(type) {
		case Token:
			ts = append(ts, p)
		case string:
			ts = append(ts, Text_(p))
		default:
			panic(errors.New(fmt.Sprintf("unsupported type %T", p)))
		}
	}
	return
}
