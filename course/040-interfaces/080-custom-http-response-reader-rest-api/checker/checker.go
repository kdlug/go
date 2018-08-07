package checker

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Results holds status codes
type Results map[string]int

// Check returns
func Check(urls []string) Results {
	res := Results{}

	for _, url := range urls {
		resp, err := http.Get(url)
		if err == nil {
			res[url] = resp.StatusCode
		}
	}
	return res
}

//   Read(p []byte) (n int, err error)
func (res Results) Read(p []byte) (n int, err error) {
	jb, err := json.Marshal(res)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	for i, v := range jb {
		p[i] = v
	}

	return len(p), io.EOF
}
