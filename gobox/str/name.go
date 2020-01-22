package str

import "strings"

func ToUp(s string) string {
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
