package node

import (
	"regexp"

	"github.com/ltto/T/gobox/str"
)

func NewNodeText(str string) *Text {
	var reg, _ = regexp.Compile("[ |\n|\t]+")
	return &Text{Text: reg.ReplaceAllString(str, " ")}
}

type Text struct {
	Text string
}

func (n *Text) Pare(m map[string]interface{}) (s string, err error) {
	conv := m["conv"].(map[string]string)
	expand := str.Expand('#', n.Text, func(s string) string {
		s2, ok := conv[s]
		if ok {
			return s2
		}
		return "#{" + s + "}"
	})
	return expand, nil
}