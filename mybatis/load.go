package mybatis

import (
	"fmt"
	"regexp"

	"github.com/beevik/etree"
	"github.com/ltto/T/mybatis/node"
)

func MainLoad() {

	Load("/Users/ltt/go/src/github.com/ltto/T/mybatis/AlbumsMapper.xml")

}

func Load(XMLPath string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(XMLPath); err != nil {
		panic(err)
	}
	element := doc.SelectElement("mapper")
	m := findSQLTPL(element, element.ChildElements())
	root := node.NewNodeRoot(m["insert"][0])
	sql, _ := PareSQL(map[string]interface{}{}, root)
	fmt.Println(sql)
}

func findSQLTPL(root *etree.Element, list []*etree.Element) map[string][]*etree.Element {
	m := make(map[string][]*etree.Element)
	for i := range list {
		elem := list[i]
		var reg, _ = regexp.Compile("[ |\n|\t]+")
		elem.SetText(reg.ReplaceAllString(elem.Text(), " "))
		if _, ok := m[elem.Tag]; !ok {
			m[elem.Tag] = root.SelectElements(elem.Tag)
		}
	}
	return m
}
