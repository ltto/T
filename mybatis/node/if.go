package node

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)


type If struct {
	Child []Node
	Test  string
}

func (n *If) pare(args map[string]interface{}) (s string, err error) {
	pareIF, err := n.pareIF(args)
	if err != nil {
		return s, err
	}
	if pareIF {
		return PareNodes(args, n.Child)
	} else {
		return "", nil
	}
}

func (n *If) pareIF(args map[string]interface{}) (bool, error) {
	parseBool, err := strconv.ParseBool(n.Test)
	if err == nil {
		return parseBool, nil
	}
	if strings.Contains(n.Test, "=") {
		split := strings.Split(n.Test, "=")
		if len(split) != 2 {
			return false, errors.New("bad if " + n.Test)
		}
		str0, err := FindStr(args, split[0])
		if err != nil {
			return false, err
		}
		str1, err := FindStr(args, split[1])
		if err != nil {
			return false, err
		}
		return fmt.Sprintf("%v", str0) == fmt.Sprintf("%v", str1), nil
	} else {
		xmlStr, err := FindStr(args, n.Test)
		if err != nil {
			return false, nil
		}
		if b, ok := xmlStr.(bool); ok {
			return b, nil
		} else {
			return false, errors.New("bad if test not bool")
		}
	}
}
