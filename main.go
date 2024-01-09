package main

import (
	"fmt"
	"net/http"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

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

// func pathHandler(w http.ResponseWriter, r *http.Request) {

// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	}
// }

func main() {
	//http.HandleFunc("/", pathHandler) // "/" path, homeHandler is the func that will process the web request
	//	http.HandleFunc("/contact", contactHandler)

	// http.Handler -> interface with the ServeHTTP method
	// http.Handle("/", http.Handler) => handle receives a Handler
	// http.HandlerFunc -> a function type that accepts same args as ServeHTTP method. also implements http.Handler
	// http.HandleFunc("/contact", contactHandler) -> receives a handlerFunc
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
