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
		panic(err)
	}
	err = json.NewDecoder(open).Decode(&Conf)
	if err != nil {
		panic(err)
	}
}

type Sys struct {
	Server struct {
		Port   int    `json:"port"`
		Static string `json:"static"`
		HTML   string `json:"html"`
		Base   string `json:"base"`
	} `json:"server"`
}
