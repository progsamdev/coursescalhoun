package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

type Template struct {
	htmlTpl *template.Template
}

func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (*Template, error) {
	tpl := template.New(patterns[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrFieldFunc": func() template.HTML {
				return `<!--csrFieldFuncPlaceholder-->`
			},
		},
	)

	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		log.Printf("parsing template %v", err)
		return &Template{}, fmt.Errorf("parse fs template: %w", err)
	}
	t := Template{
		htmlTpl: tpl,
	}
	return &t, nil
}

/*func Parse(filePath string) (*Template, error) {
	tpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Printf("parsing template %v", err)
		return &Template{}, fmt.Errorf("parsing template: %w", err)
	}
	t := Template{
		htmlTpl: tpl,
	}
	return &t, nil
}*/

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		return
	}

	tpl = tpl.Funcs(
		template.FuncMap{
			"csrFieldFunc": func() template.HTML {
				return csrf.TemplateField(r)
			},
		},
	)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tpl.Execute(w, data)
	if err != nil {
		log.Printf("parsing template %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
