package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	// httptest has two methods: NewRequest and NewRecorder, which help simplify the process of running tests against your handlers.
	// NewRequest mocks a request that would be used to serve your handler.
	req := httptest.NewRequest("GET", "http://google,com", nil)
	// NewRecorder is a drop-in replacement for ResponseWriter and is used to process and compare the response with the expected output.
	w := httptest.NewRecorder() // ResponseWriter interface - records all the responses from the handler:
	handler(w, req)
	resp := w.Result()
	fmt.Println(resp.StatusCode)
}
