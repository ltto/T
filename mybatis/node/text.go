package node

import (
	"github.com/ltto/T/gobox/str"
)

type Text struct {
	Text string
}

func (n *Text) pare(args map[string]interface{}) (s string, err error) {
	temp := args["_temp"].(map[string]string)
	expand := str.Expand('#', n.Text, func(s string) string {
		s2, ok := temp[s]
		if ok {
			return s2
		}
		return "#{" + s + "}"
	})
	return expand, nil
}
