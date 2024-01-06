package main

import (
	"fmt"
	"net/http"
)

const PAGE_NOT_FOUND = "PAGE NOT FOUND"

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8") //specify the type of content so browser can handle it properly // set instead of add so we override any other setting
	fmt.Fprint(w, "<h1> Welcome to my cool site! </h1>")
	//w is responsable for the response
	//r is the request
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> Contact page </h1> <p> To get in touch, email me at <a href=\"mailto:samuel.msbr@gmail.com\">samuel.msbr@gmail.com</a>.")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		msg := PAGE_NOT_FOUND
		errorHandler(w, r, msg, http.StatusNotFound)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, msg string, status int) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "%s : %v", msg, status)
}

func main() {
	http.HandleFunc("/", pathHandler) // "/" path, homeHandler is the func that will process the web request
	//	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil) //build server in port 3000
}
