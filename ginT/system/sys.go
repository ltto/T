package system

import (
	"encoding/json"
	"os"
)

var (
	Conf Sys
)

func init() {
	open, err := os.Open("conf.json")
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
		Dest   string `json:"dest"`
		HTML   string `json:"html"`
		Base   string `json:"base"`
	} `json:"server"`
}
