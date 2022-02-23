package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime/debug"
	"strings"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ltto/T/gobox/gen"
	_ "embed"
	"io/ioutil"
)

type L struct {
	Deep  int
	Close bool
}

func (l L) Print(v ...interface{}) {
	if l.Close {
		return
	}
	space := l.deep()
	var vv = []interface{}{space}
	fmt.Println(append(vv, v...)...)
}
func (l L) Printf(format string, v ...interface{}) {
	if l.Close {
		return
	}
	space := l.deep()
	fmt.Printf(space+" "+format+"\n", v...)
}

func (l L) deep() string {
	buffer := bytes.NewBuffer(debug.Stack())
	var leve string
	for i := 0; i < l.Deep*2; i++ {
		ss, _ := buffer.ReadString('\n')
		if i == l.Deep*2-3 {
			leve = ss[strings.LastIndex(ss, ".")+1 : strings.LastIndex(ss, "(")]
		}
	}
	s, _ := buffer.ReadString('\n')
	split := strings.Split(s, " ")
	space := strings.TrimSpace(split[0])
	f := fmt.Sprintf("[%-6v] ", leve)
	return f + space
}

//go:embed resources/index.html
var indexHTML []byte

//go:embed resources/jq.js
var jqJS []byte

func main() {
	// Set logger
	//l := log.New(log.Writer(), log.Prefix(), log.Flags()|log.Llongfile)
	l := L{Deep: 5, Close: true}
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	// Create astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:           "BestApp",
		BaseDirectoryPath: path.Join(home, ".astilectron"),
		DataDirectoryPath: path.Join(home, ".astilectron"),
	})
	if err != nil {
		panic(err)
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		panic(err)
	}

	// New window
	var w *astilectron.Window
	o := &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(980),
		Width:  astikit.IntPtr(1440),
		WebPreferences: &astilectron.WebPreferences{
			NodeIntegrationInWorker: astikit.BoolPtr(true),
		},
	}
	html := path.Join(home, ".astilectron", "resources", "index.html")
	jq := path.Join(home, ".astilectron", "resources", "jq.js")

	_ = os.MkdirAll(path.Dir(html), 0777)
	if err = ioutil.WriteFile(html, indexHTML, 0777); err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile(jq, jqJS, 0777); err != nil {
		panic(err)
	}

	if w, err = a.NewWindow(html, o); err != nil {
		panic(err)
	}
	// Create windows
	if err = w.Create(); err != nil {
		panic(err)
	}
	// This will listen to messages sent by Javascript

	w.OnMessage(func(m *astilectron.EventMessage) interface{} {
		var s string
		fmt.Println(s)
		if err := m.Unmarshal(&s); err != nil {
			return rep{-1, err.Error()}
		}
		var msg Msg
		if err := json.Unmarshal([]byte(s), &msg); err != nil {
			return rep{-1, err.Error()}
		}
		f, err := msg.f()
		if err != nil {
			return rep{-1, err.Error()}
		}
		return rep{0, f}
	})

	// Blocking pattern
	a.Wait()
}

type rep struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Msg struct {
	F    string    `json:"f"`
	Data gen.Param `json:"data"`
}

func (m Msg) f() (interface{}, error) {
	switch m.F {
	case "scan":
		param := m.Data
		if param.Regexp == "" {
			param.Regexp = ".*"
		}
		compile, err := regexp.Compile(param.Regexp)
		if err != nil {
			return nil, err
		}
		if !(param.Password == "") {
			param.Password = ":" + param.Password
		}
		url := fmt.Sprintf("%v%v@tcp(%v:%v)/%v?charset=utf8mb4&collation=utf8mb4_bin&loc=Local&parseTime=true",
			param.User,
			param.Password,
			param.Host,
			param.Port,
			param.Database,
		)
		db, err := sql.Open("mysql", url)
		if err != nil {
			return nil, err
		}
		rows, err := db.Query("SHOW TABLES ")
		if err != nil {
			return nil, err
		}
		var tables []string
		var tab string
		for rows.Next() {
			if err := rows.Scan(&tab); err != nil {
				return nil, err
			}
			if compile.Match([]byte(tab)) {
				tables = append(tables, tab)
			}
		}
		return tables, nil

	case "gen":
		result, err := gen.ScanTable(&m.Data)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, errors.New("fname err")
}
