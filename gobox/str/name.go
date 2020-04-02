package str

import (
	"bytes"
	"strings"
)

func ToCamel(s string) string {
	newStr := ""
	Up := false
	for _, v := range s {
		if v == '_' {
			Up = true
			continue
		}
		strCell := ""
		if Up && v != '_' {
			strCell = strings.ToUpper(string(v))
			Up = false
		} else {
			strCell = string(v)
		}
		newStr += strCell
	}
	return strings.ToUpper(newStr[0:1]) + newStr[1:]
}
func Underline(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		t = append(t, 'X')
		i++
	}
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIIUpper(s[i+1]) {
			continue
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}

		if isASCIIUpper(c) {
			c ^= ' '
		}
		t = append(t, c)

		for i+1 < len(s) && isASCIIUpper(s[i+1]) {
			i++
			t = append(t, '_')
			t = append(t, bytes.ToLower([]byte{s[i]})[0])
		}
	}
	return string(t)
}
func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}


func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}