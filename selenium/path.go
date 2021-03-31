package selenium

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func FindPATH(file string) string {
	pathEnvs := strings.Split(os.Getenv("PATH"), string(filepath.ListSeparator))
	var suffix = []string{""}
	if runtime.GOOS == "windows" {
		suffix = []string{"", ".cmd", ".exe", ".com", ".bat"}
	}
	for _, env := range pathEnvs {
		for _, suf := range suffix {
			var pah = ""
			if path.Ext(file) == "" {
				pah = path.Join(env, file+suf)
			} else {
				path.Join(env, file)
			}
			if Exist(pah) {
				return pah
			}
		}
	}
	return file
}
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	} else {
		return !os.IsNotExist(err)
	}
}
