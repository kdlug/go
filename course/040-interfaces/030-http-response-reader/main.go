// reader and writer
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

	// headers etc, without body
	fmt.Println(resp)

	// body is a field of response struct, and it has a type interface
	// type Response struct {
	//     Status     string // e.g. "200 OK"
	//     StatusCode int    // e.g. 200
	//     Proto      string // e.g. "HTTP/1.0"
	//     ProtoMajor int    // e.g. 1
	//     ProtoMinor int    // e.g. 0
	//     Header Header
	//     Body io.ReadCloser
	//     ContentLength int64
	//     TransferEncoding []string
	//     Close bool
	//     Uncompressed bool
	//     Trailer Header
	//     Request *Request
	//     TLS *tls.ConnectionState
	// }

	// It gives us a little bit more flexibility. We can put to Body field any type until it fulfills ReadCloserInterface
	// type ReadCloser interface {
	//     Reader
	//     Closer
	// }

	// bs := []byte{} // it won't work because the slice is empty
	bs := make([]byte, 99999) // we have to pass a slice with some capacity to the reader
	resp.Body.Read(bs)

	//fmt.Println(string(bs))

	// #2 we can use a helper
	// writer interface
	io.Copy(os.Stdout, resp.Body) //(writer interface, reader interface)
}
