package main

import (
	"fmt"
	"net/http"
)

// handler receives all incomming HTTP connections
// http.ResponseWriter is a text/html response
// http.Request contains all information about this request
// Fex:
// GET: request.URL.Query().Get("parameter")
// POST: request.FormValue("name")
func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Requested URL: %s\n", request.URL.Path)
}

func main() {
	// register a handler
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
