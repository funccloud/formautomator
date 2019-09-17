package main

import (
	"fmt"
	fa "formautomator"
	"io/ioutil"
	"os"

	"github.com/crgimenes/goconfig"
)

type config struct {
	Filename string `json:"filename,omitempty" cfg:"f" cfgRequired:"true" cfgDefault:"-"`
}

func main() {
	cfg := config{}
	err := goconfig.Parse(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	var b []byte
	if cfg.Filename == "-" {
		b, err = ioutil.ReadAll(os.Stdin)
	} else {
		b, err = ioutil.ReadFile(cfg.Filename)
	}
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
