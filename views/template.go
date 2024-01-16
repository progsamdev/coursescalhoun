package views

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Template struct {
	htmlTpl *template.Template
}

func Parse(filePath string) (Template, error) {

	tpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Printf("parsing template %v", err)
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	t := Template{
		htmlTpl: tpl,
	}
	return t, nil
}

func (t Template) execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("parsing template %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
