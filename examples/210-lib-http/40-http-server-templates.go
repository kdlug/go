// html/template package supports HTML templates
// automatically escapes data which means there is no need to worry about about XSS attacks
// To access the data in a template the top most variable is access by {{.}}.
// The dot inside the curly braces is called the pipeline and the root element of the data.
// https://golang.org/pkg/text/template/#hdr-Actions
// localhost:8090/blog/
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Title   string
	Visited int
}

type PostListPage struct {
	PageTitle string
	Posts     []Post
}

func handler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintf(response, "Requested URL: %s\n", request.URL.Path)
	fmt.Fprintf(response, "Vars: %s\n", vars)
}

func postList(response http.ResponseWriter, request *http.Request) {
	// parse template
	tmpl := template.Must(template.ParseFiles("layout.html"))

	// data
	data := PostListPage{
		PageTitle: "Blog Post List",
		Posts: []Post{
			{Title: "Post 1", Visited: 0},
			{Title: "Post 2", Visited: 100},
			{Title: "Post 3", Visited: 23},
		},
	}

	// Execute a Template in a Request Handler
	tmpl.Execute(response, data)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/blog/{title}/page/{page}", handler)
	// post list
	router.HandleFunc("/blog/", postList)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8090", router)
}
