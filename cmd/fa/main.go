package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	fa "code.funccloud.dev/formautomator"
	"github.com/crgimenes/goconfig"
)

type config struct {
	Filename  string `json:"filename,omitempty" cfg:"f" cfgRequired:"true" cfgDefault:"-"`
	Templates string `json:"templates,omitempty" cfg:"t" cfgRequired:"true"`
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
	templates, err := filepath.Glob(path.Join(cfg.Templates, "*.html"))
	if err != nil {
		fmt.Println(err)
	}
	s, err := fa.CreateForm(b, templates)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
