// get keys and list k/v pairs
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

	// PUT a few key/value pairs
	pairs := api.KVPairs{
		&api.KVPair{Key: "service-name/domain1", Value: []byte("www.example.com")},
		&api.KVPair{Key: "service-name/domain2", Value: []byte("www.google.com")},
		&api.KVPair{Key: "service-name/domain3", Value: []byte("www.amazon.com")},
	}

	for _, pair := range pairs {
		_, err = kv.Put(pair, nil)
		if err != nil {
			panic(err)
		}
	}

	// Get keys
	keys, _, err := kv.Keys("service-name", "|", nil)
	fmt.Println(keys)

	// Get a list
	list, _, err := kv.List("service-name", nil)

	for _, pair := range list {
		fmt.Println(string(pair.Key), "=>", string(pair.Value))
	}

	// Cleanup
	for _, pair := range pairs {
		_, err = kv.Delete(pair.Key, nil)
		if err != nil {
			panic(err)
		}
	}
}
