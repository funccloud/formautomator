package main

import (
	"fmt"
	fa "formautomator"
	"io/ioutil"

	"github.com/crgimenes/goconfig"
)

type config struct {
	Filename string `json:"filename,omitempty" cfg:"f" cfgRequired:"true"`
}

func main() {
	cfg := config{}
	err := goconfig.Parse(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := ioutil.ReadFile(cfg.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	s, err := fa.CreateForm(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
