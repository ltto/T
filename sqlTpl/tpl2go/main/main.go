package main

import (
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var file string
var varName string
var pkg string
var dir string
var t string

func main() {
	flag.StringVar(&file, "file", "", "tpl dest dir")
	flag.StringVar(&varName, "name", "TPLS", "map var Name")
	flag.StringVar(&pkg, "pkg", "s", "target file package")
	flag.StringVar(&dir, "dir", "tpl", "dir")
	flag.StringVar(&t, "t", "./bytes_tpl.go", "target file")
	flag.Parse()
	dest := path.Join(file, "../", dir)
	file := path.Join(file, "../", t)
	dir, err := ioutil.ReadDir(dest)
	if err != nil {
		panic(err)
	}
	m := make(map[string][]byte, len(dir))
	for _, info := range dir {
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".gohtml") {
			continue
		}
		bytes, err := ioutil.ReadFile(path.Join(dest, info.Name()))
		if err != nil {
			panic(err)
		}
		m[info.Name()] = bytes
	}
	s := "package " + pkg + "\n" +
		"func init(){\n" +
		wireMAP(m) +
		"}\n"
	source, err := format.Source([]byte(s))
	if err != nil {
		panic(err)
	}

	if err = os.Remove(file); err != nil {
		panic(err)
	}

	if err = ioutil.WriteFile(file, source, 0666); err != nil {
		panic(err)
	}

}
func wireMAP(m map[string][]byte) string {
	str := varName + " = map[string][]byte{\n"
	for s, bytes := range m {
		str += "\"" + s + "\": {" + Bytes(bytes) + "},\n"
	}
	str += "}"
	return str
}
func Bytes(b []byte) string {
	str := ""
	for _, v := range b {
		str += fmt.Sprintf(" %d,", v)
	}
	return str[:len(str)-1]
}
