package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		log.Printf("parsing template %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("there was an error executing template  %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusOK)
	// fmt.Fprint(w)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> Contact page </h1> <p> To get in touch, email me at <a href=\"mailto:samuel.msbr@gmail.com\">samuel.msbr@gmail.com</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	color := chi.URLParam(r, "color")

	html := `<h1 `
	if color != "" {
		html += `style="background-color: ` + strings.Trim(color, "{}") + `;"`
	}
	html += `> Welcome to my site </h1> 
	
	<h2> Q: How can I contact you? </h2>
	<p> A: You can reach me via email at <a href="mailto:samuel.msbr@gmail.com">samuel.msbr@gmail.com</a>.</p>

	<h2> Q: What services do you provide? </h2>
	<p> A: Currently, I do not provide specific services through this platform. If you have any inquiries, feel free to contact me via email.</p>

	<h2> Q: Can I collaborate with you on a project? </h2>
	<p> A: I'm open to collaboration opportunities. Please reach out via email to discuss further details.</p>

	<h2> Q: How long does it take to receive a response? </h2>
	<p> A: I strive to respond to emails as promptly as possible. However, response times may vary based on workload and other commitments.</p>

	<h2> Q: Do you have a social media? </h2>
	<p> A: At the moment, I don't have active social media profiles. Email is the best way to get in touch.</p>`

	fmt.Fprint(w, html)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/faq/{color}", faqHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
