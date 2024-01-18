package controllers

import (
	"net/http"

	"github.com/progsamdev/coursescalhoun/views"
)

func StaticHandler(tpl *views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
