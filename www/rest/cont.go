package rest

import "strings"

const (
	httpRedirect = "redirect:"
	httpFile     = "file:"
	httpHtml     = "html:"
)

func ReturnRedirect(path string) string {
	return httpRedirect + path
}
func ReturnFile(path string) string {
	return httpFile + path
}
func ReturnHtml(path string) string {
	return httpHtml + path
}
func Redirect(s string) (ss string, b bool) {
	b = strings.HasPrefix(s, httpRedirect)
	if b {
		ss = s[len(httpRedirect):]
	}
	return
}
func File(s string) (ss string, b bool) {
	b = strings.HasPrefix(s, httpFile)
	if b {
		ss = s[len(httpFile):]
	}
	return
}
func Html(s string) (ss string, b bool) {
	b = strings.HasPrefix(s, httpHtml)
	if b {
		ss = s[len(httpHtml):]
	}
	return
}
