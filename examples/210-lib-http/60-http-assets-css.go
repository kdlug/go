package main

import (
	"net/http"
)

func main() {
	// To serve static assets like JavaScript, CSS and images, we use the inbuilt http.FileServer and point it to a url path
	fs := http.FileServer(http.Dir("assets/")) // directory on the hard drive

	// point a URL path to the file server
	// curl -s http://localhost:8080/static/css/styles.css
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8090", nil)
}
