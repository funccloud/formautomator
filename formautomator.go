package formautomator

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/yosssi/gohtml"
)

// Field properties
type Field struct {
	Name        string `json:"name,omitempty"`
	Label       string `json:"label,omitempty"`
	Class       string `json:"class,omitempty"`
	Type        string `json:"type,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	Value       string `json:"value,omitempty"`
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
		name="{{name}}"
		value="{{value}}">
	`
)

// CreateForm generate an HTML form
func CreateForm(j json.RawMessage, templates []string) (string, error) {
	t := make(map[string]*template.Template)
	for _, v := range templates {
		basename := filepath.Base(v)
		name := strings.TrimSuffix(basename, filepath.Ext(v))
		tAux, err := template.New(name).ParseFiles(v)
		if err != nil {
			return "", err
		}
		t[name] = tAux
	}

	fields := []Field{}
	err := json.Unmarshal(j, &fields)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	s := ""
	for _, f := range fields {
		if f.Class == "" {
			f.Class = "form-control"
		}
		if f.Type == "" {
			f.Type = "text"
		}

		tpl := t[f.Type]
		err = tpl.Execute(buf, f)
		if err != nil {
			return "", err
		}
		s += buf.String()
	}

	formStru := struct {
		Fields string
	}{
		Fields: s,
	}
	tpl := t["form"]
	err = tpl.Execute(buf, formStru)
	if err != nil {
		return "", err
	}
	s = buf.String()
	s = strings.TrimSpace(s)
	s = strings.Join(strings.Fields(s), " ")
	s = gohtml.Format(s)
	return s, nil
}
