// gorilla/mux - powerfull router and dispatcher for go
// go get -u github.com/gorilla/mux
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(response http.ResponseWriter, request *http.Request) {
	// 	mux.Vars(r) which takes the http.Request as parameter and returns a map of route variables
	vars := mux.Vars(request)
	fmt.Fprintf(response, "Requested URL: %s\n", request.URL.Path)
	fmt.Fprintf(response, "Vars: %s\n", vars)
}

func main() {
	// create a router
	// it will receive all HTTP connections and pass it on to the request handlers you will register on it
	router := mux.NewRouter()

	// now register request handlers with one difference: use router.HandleFunc(...) instead of htt.HandleFunc(...)
	// f.ex we want to handle the following url, which is dynamic (title and page number)
	// /blog/whats-new/page/2
	router.HandleFunc("/blog/{title}/page/{page}", handler)

	// we can restrict handler to use a specific HTTP method
	// router.HandleFunc("/blog/{title}", CreatePost).Methods("POST")
	// router.HandleFunc("/blog/{title}", ReadPost).Methods("GET")
	// router.HandleFunc("/blog/{title}", UpdatePost).Methods("PUT")
	// router.HandleFunc("/blog/{title}", DeletePost).Methods("DELETE")
	//
	// or specific domains
	// router.HandleFunc("/blog/{title}", handler).Host("www.myblog.com")
	//
	// or prefixes
	// blogrouter := r.PathPrefix("/blog").Subrouter()
	// blogrouter.HandleFunc("/", AllPosts)
	// blogrouter.HandleFunc("/{title}", GetPost)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// the second parameter is the router, nil means the default router
	// let's change it to gorilla/mux
	http.ListenAndServe(":8090", router)
}
