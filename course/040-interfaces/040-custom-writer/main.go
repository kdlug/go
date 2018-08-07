package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// custom writer
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// custom writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// logWriter implements writer interface
// (logWriter) -> a shortcut because we don't use variable inside the function
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote:", len(bs))
	return len(bs), nil
}
