package www

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
