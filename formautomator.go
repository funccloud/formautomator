package formautomator

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/yosssi/gohtml"
)

type Metadata struct {
	ForList string `json:"form_list,omitempty"`
	Index   string `json:"index,omitempty"`
	Label   string `json:"label,omitempty"`
	Var     string `json:"var,omitempty"`
}

// Field properties
type Field struct {
	Name        string   `json:"name,omitempty"`
	Label       string   `json:"label,omitempty"`
	Class       string   `json:"class,omitempty"`
	Type        string   `json:"type,omitempty"`
	Placeholder string   `json:"placeholder,omitempty"`
	Value       string   `json:"value,omitempty"`
	Metadata    Metadata `json:"metadata,omitempty"`
}

type Form struct {
	Fields  []Field `json:"fields,omitempty"`
	Method  string  `json:"method,omitempty"`
	Action  string  `json:"action,omitempty"`
	EncType string  `json:"enctype,omitempty"`
}

// CreateForm generate an HTML form
func CreateForm(j json.RawMessage, templates []string) (string, error) {
	t := make(map[string]*template.Template)
	for _, v := range templates {
		basename := filepath.Base(v)
		name := strings.TrimSuffix(basename, filepath.Ext(v))
		b, err := ioutil.ReadFile(v)
		if err != nil {
			return "", err
		}
		tAux, err := template.New(name).Delims("[[", "]]").Parse(string(b))
		if err != nil {
			return "", err
		}
		t[name] = tAux
	}
	f := Form{}
	err := json.Unmarshal(j, &f)
	if err != nil {
		return "", err
	}
	s := ""
	for _, f := range f.Fields {
		buf := &bytes.Buffer{}
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
		Fields  string `json:"fields,omitempty"`
		Method  string `json:"method,omitempty"`
		Action  string `json:"action,omitempty"`
		EncType string `json:"enctype,omitempty"`
	}{
		Fields:  s,
		Method:  f.Method,
		Action:  f.Action,
		EncType: f.EncType,
	}
	buf := &bytes.Buffer{}
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
