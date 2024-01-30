package main

//
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Star Wars IV", Description: "Excellent Movie", Price: 10.00},
		Article{Title: "Moby Dick", Description: "Well-known book", Price: 25.00},
	}

	fmt.Println(w, "Get All Articles Endpoint")
	// encode to json
	json.NewEncoder(w).Encode(articles)

}

func handleRequests() {
	// register function
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", getAllArticles)
	// Your http.ListenAndServe function passes a nil value for the http.Handler parameter.
	// This tells the ListenAndServe function that you want to use the default server multiplexer
	// and not the one you’ve set up.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
