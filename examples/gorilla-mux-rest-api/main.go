package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func postArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Article Endpoint")
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
	// define new router
	router := mux.NewRouter().StrictSlash(true)
	// register function
	// replace http with router
	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", getAllArticles).Methods("GET") // thanks to gorilla we can specify HTTP methods to the same path
	router.HandleFunc("/articles", postArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router)) // pass created router
}

func main() {
	handleRequests()
}
