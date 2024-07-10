// We are defining an interfase using a value semantics
package main

import "fmt"

// reader is an interface that defines the act of reading data
type reader interface {
	read(b []byte) (int, error)
}

// file defines a system file
type file struct {
	name string
}

// read implements the reader interface for a files
func (file) read(b []byte) (int, error) {
	s := "<rss><channel><title> GO Programming</title></channel></rss>"
	copy(b, s)

	return len(s), nil
}

// pipe defines a named pipe network connection.
type pipe struct {
	name string
}

// read implements reader interface for a network connection
func (pipe) read(b []byte) (int, error) {
	s := `{name: "pablo", title: "painter"}`
	copy(b, s)
	return len(s), nil
}

func retrieve(r reader) error {
	data := make([]byte, 100)

	len, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}

func main() {
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	retrieve(f)
	retrieve(p)
}
