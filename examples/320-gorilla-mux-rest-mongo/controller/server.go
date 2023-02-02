package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/kdlug/restapi/controller/blog"
)

func HandleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/api/posts", controller.GetAllPosts).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router)) // pass created router
}
