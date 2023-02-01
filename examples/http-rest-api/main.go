package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}