// go get github.com/hashicorp/consul/api
// http://techblog.zeomega.com/devops/golang/2015/06/09/consul-kv-api-in-golang.html
package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	kv := client.KV()

	// DELETE foo key
	_, err = kv.Delete("foo", nil)
	if err != nil {
		fmt.Println(err)
	}

	// DELETE service-name folder
	_, err = kv.DeleteTree("service-name", nil)
	if err != nil {
		fmt.Println(err)
	}
}
