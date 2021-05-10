package node

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type IF struct {
	Child []Node
	Test  string
}

func (n *IF) pare(args map[string]interface{}) (s string, err error) {
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

func (n *IF) pareIF(args map[string]interface{}) (bool, error) {
	parseBool, err := strconv.ParseBool(n.Test)
	if err == nil {
		return parseBool, nil
	}
	s1, s2, fuc, err := doSplit(n.Test, args)
	if err != nil {
		return false, err
	}
	if fuc != 0 {
		if fuc == '=' {
			if fmt.Sprintf("%v", s1) == fmt.Sprintf("%v", s2) {
				return true, nil
			}
		}
		return doCmp(fmt.Sprintf("%v", s1), fmt.Sprintf("%v", s2), fuc)
	} else {
		xmlStr, err := FindStr(args, n.Test)
		if err != nil {
			return false, nil
		}
		if b, ok := xmlStr.(bool); ok {
			return b, nil
		} else {
			if b, err = strconv.ParseBool(fmt.Sprintf("%v", xmlStr)); err != nil {
				return false, errors.New("bad if test not bool")
			} else {
				return b, nil
			}
		}
	}
}

func doCmp(s1, s2 string, fuc rune) (ok bool, err error) {
	var f1, f2 float64
	if f1, err = strconv.ParseFloat(s1, 10); err != nil {
		return false, err
	}
	if f2, err = strconv.ParseFloat(s2, 10); err != nil {
		return false, err
	}
	cmp := big.NewFloat(f1).Cmp(big.NewFloat(f2))
	switch fuc {
	case '=':
		return cmp == 0, nil
	case '>':
		return cmp > 0, nil
	case '<':
		return cmp < 0, nil
	default:
		return false, errors.New("cmp func not found")
	}
}
func doSplit(test string, args map[string]interface{}) (s1, s2 interface{}, fuc rune, err error) {
	fuc = 0
	if strings.Contains(test, "=") {
		fuc = '='
	} else if strings.Contains(test, ">") {
		fuc = '>'
	} else if strings.Contains(test, "<") {
		fuc = '<'
	} else {
		return
	}
	split := strings.Split(test, string(fuc))
	if len(split) != 2 {
		err = errors.New("bad if " + test)
		return
	}
	if s1, err = FindStr(args, split[0]); err != nil {
		return
	}
	if s2, err = FindStr(args, split[1]); err != nil {
		return
	}
	return
}
