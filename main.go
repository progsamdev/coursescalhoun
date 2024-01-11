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
	case "/faq":
		faqHandler(w, r)
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

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<h1> Frequently Asked Questions </h1>
		
		<h2> Q: How can I contact you? </h2>
		<p> A: You can reach me via email at <a href="mailto:samuel.msbr@gmail.com">samuel.msbr@gmail.com</a>.</p>

		<h2> Q: What services do you provide? </h2>
		<p> A: Currently, I do not provide specific services through this platform. If you have any inquiries, feel free to contact me via email.</p>

		<h2> Q: Can I collaborate with you on a project? </h2>
		<p> A: I'm open to collaboration opportunities. Please reach out via email to discuss further details.</p>

		<h2> Q: How long does it take to receive a response? </h2>
		<p> A: I strive to respond to emails as promptly as possible. However, response times may vary based on workload and other commitments.</p>

		<h2> Q: Do you have a social media presence? </h2>
		<p> A: At the moment, I don't have active social media profiles. Email is the best way to get in touch.</p>
	`)
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
