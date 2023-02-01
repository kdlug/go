package main

// The middleware performs some specific function on the HTTP request or response
// at a specific stage in the HTTP pipeline before or after the user defined controller.
// Middleware is a design pattern to eloquently add cross cutting concerns
// like logging, handling authentication, or gzip compression without having many code contact points.
// In other languages it can be called middleware or filter
import (
	"fmt"
	"log"
	"net/http"
)

// Adds X-App-Version to the response header
func appVersionMiddleware(handler http.HandlerFunc, version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-App-Version", version)
		handler.ServeHTTP(w, r)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint")
}

func handleRequests() {
	// register function
	http.HandleFunc("/", appVersionMiddleware(homePage, "v1"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
