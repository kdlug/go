package main

import (
	"fmt"
	"net/http"
)

func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Requested URL: %s\n", request.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)

	// To serve static assets like JavaScript, CSS and images, we use the inbuilt http.FileServer and point it to a url path
	fs := http.FileServer(http.Dir("static/"))

	// point a URL path to the file sercer
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8090", nil)
}
