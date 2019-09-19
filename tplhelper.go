package formautomator

import (
	"html/template"
	"reflect"
)

var TemplateFunctions = template.FuncMap{
	"in": in,
}

func in(a []interface{}, b interface{}) bool {
	for _, v := range a {
		if reflect.DeepEqual(v, b) {
			return true
		}
	}
	return false
}
