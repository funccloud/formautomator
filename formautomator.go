package formautomator

import (
	"encoding/json"
	"strings"

	"github.com/yosssi/gohtml"
)

// Field properties
type Field struct {
	Name        string `json:"name,omitempty"`
	Label       string `json:"label,omitempty"`
	Class       string `json:"class,omitempty"`
	Type        string `json:"type,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
}

const (
	tForm = `
	<form>{{form}}</form>
	`
	tField = `
	<label for="{{name}}">
		{{label}}
	</label>
	<input
		class="{{class}}"
		type="{{type}}"
		placeholder="{{placeholder}}"
		name="{{name}}">
	`
)

// CreateForm generate an HTML form
func CreateForm(j json.RawMessage) (string, error) {
	fields := []Field{}
	err := json.Unmarshal(j, &fields)
	if err != nil {
		return "", err
	}
	s := ""
	for _, f := range fields {
		if f.Class == "" {
			f.Class = "form-control"
		}
		if f.Type == "" {
			f.Type = "text"
		}
		a := strings.ReplaceAll(tField, "{{name}}", f.Name)
		a = strings.ReplaceAll(a, "{{label}}", f.Label)
		a = strings.ReplaceAll(a, "{{type}}", f.Type)
		a = strings.ReplaceAll(a, "{{placeholder}}", f.Placeholder)
		a = strings.ReplaceAll(a, "{{class}}", f.Class)
		s += a
	}
	s = strings.ReplaceAll(tForm, "{{form}}", s)
	s = strings.TrimSpace(s)
	s = strings.Join(strings.Fields(s), " ")
	s = gohtml.Format(s)
	return s, nil
}
