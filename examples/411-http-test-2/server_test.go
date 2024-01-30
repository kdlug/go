package main

// https://speedscale.com/testing-golang-with-httptest/
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler(t *testing.T) {
	expected := "Hello Amigo"

	req := httptest.NewRequest(http.MethodGet, "/hello?name=Amigo", nil) // creates a mock request to /hello with a name parameter
	w := httptest.NewRecorder()

	RequestHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected %s but got %v", expected, string(data))
	}
}

func TestRequestHandlerJson(t *testing.T) {
	expected := "Hello Amigo"

	res := callHelloJsonEndpoint()
	defer res.Body.Close()

	// If we read from res.Body which is io.ReadCloser we gets Error: EOF
	//body, _ := io.ReadAll(res.Body)
	//fmt.Printf("response %s", body)

	jsonResponse := new(MyResponse)

	// data, err := io.ReadAll(res.Body)
	// we can also decode to object like this (useful for json responses)
	err := json.NewDecoder(res.Body).Decode(jsonResponse)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	fmt.Printf("response %v", jsonResponse)
	if jsonResponse.Message != expected {
		t.Errorf("Expected %s but got %s", expected, jsonResponse.Message)
	}

}
}

}




func callHelloEndpoint() *http.Response {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Amigo", nil) // creates a mock request to /hello with a name parameter
	w := httptest.NewRecorder()

	RequestHandler(w, req)
	res := w.Result()

	return res
}

func callHelloJsonEndpoint() *http.Response {
	req := httptest.NewRequest(http.MethodGet, "/hello_json?name=Amigo", nil) // creates a mock request to /hello with a name parameter
	w := httptest.NewRecorder()

	RequestHandler(w, req)
	res := w.Result()

	return res
}
