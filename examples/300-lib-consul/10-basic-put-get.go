// http://techblog.zeomega.com/devops/golang/2015/06/09/consul-kv-api-in-golang.html
package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {
	// Consul
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "foo", Value: []byte("test")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("foo", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(pair.Key), "=>", string(pair.Value))
}
