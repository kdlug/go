package main

// http://localhost:3030/hello?name=pablo
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}

	name := query.Get("name")

	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing name")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %s", name)
}

type MyResponse struct {
	Message string `json:"message"`
}

func RequestHandlerJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query, err := url.ParseQuery(r.URL.RawQuery)
	myResponse := new(MyResponse)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		myResponse.Message = "Bad Request"
		json.NewEncoder(w).Encode(myResponse)
		return
	}

	name := query.Get("name")

	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		myResponse.Message = "Bad Request"
		json.NewEncoder(w).Encode(myResponse)
		return
	}

	w.WriteHeader(http.StatusOK)

	myResponse.Message = "Hello " + name

	json.NewEncoder(w).Encode(myResponse)
}

func main() {
	http.HandleFunc("/hello", RequestHandler)
	http.HandleFunc("/hello_json", RequestHandlerJson)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
