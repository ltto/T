package conf

import (
	"encoding/json"
	"flag"
	"os"
)

var (
	Conf Sys
)

func init() {
	s := flag.String("c", "conf.json", "config file")
	open, err := os.Open(*s)
	if err != nil {
		Conf.Server.HostPort = ":8080"
		return
	}
	err = json.NewDecoder(open).Decode(&Conf)
	if err != nil {
		Conf.Server.HostPort = ":8080"
	}
}

type Sys struct {
	Server struct {
		HostPort    string `json:"host_port"`
		StaticLocal string `json:"static_local"`
		StaticPath  string `json:"static_path"`
		Base        string `json:"base"`
	} `json:"server"`
}

func (s Sys) GetStaticPath() string {
	if s.Server.StaticPath == "" {
		return "static"
	}
	return s.Server.StaticPath
}

func (s Sys) GetStaticLocal() string {
	if s.Server.StaticLocal == "" {
		return "static"
	}
	return s.Server.StaticLocal
}
