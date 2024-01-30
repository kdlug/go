package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var port string = "8000"
var host string = "127.0.0.1"

type requestLoggerHandler struct{}

func (h requestLoggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var bodyBytes []byte
	var err error
	fmt.Printf("Request %s %s %s \n", r.Method, r.URL.Host, r.URL.Path)
	if r.Body != nil {
		bodyBytes, err = io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			fmt.Printf("Body reading error: %v", err)
			return
		}
	}

	fmt.Printf("Headers: %+v\n", r.Header)

	if len(bodyBytes) > 0 {
		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
			fmt.Printf("JSON parse error: %v", err)
			return
		}
		fmt.Println(string(prettyJSON.Bytes()))
	} else {
		fmt.Printf("Body: No Body Supplied\n")
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Printf("provide server port\n")
		fmt.Printf("%s port_number\n", filepath.Base(os.Args[0]))
		return
	}

	if len(args) == 1 {
		if _, err := strconv.Atoi(args[0]); err != nil {
			fmt.Printf("provide correct port number\n")
			return
		}

		port = args[0]
	}

	fmt.Printf("Starting request echo server on port %v\n", port)
	addr := fmt.Sprintf("%v:%v", host, port)
	err := http.ListenAndServe(addr, requestLoggerHandler{})
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
