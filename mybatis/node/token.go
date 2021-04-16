package node

//type Pare interface {
//	PareChild(es []Token) []Node
//	NewNodeIf(token Token) *IF
//	NewNodeForeach(token Token) *Foreach
//	NewNodeInclude(refId string) *Include
//	NewNodeText(data string) *Text
//}

type Token interface {
	Tag() string
	Data() string
	Child() []Node
	Attr(key string) string
}
