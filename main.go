package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8") //specify the type of content so browser can handle it properly // set instead of add so we override any other setting
	fmt.Fprint(w, "<h1> Welcome to my cool site! </h1>")
	//w is responsable for the response
	//r is the request
}

func main() {
	http.HandleFunc("/", handlerFunc) // "/" path, handlerFunc is the func that will process the web request
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil) //build server in port 3000
}
