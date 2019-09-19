package formautomator

import (
	"html/template"
)

var TemplateFunctions = template.FuncMap{
	"in": in,
}

func in(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}
