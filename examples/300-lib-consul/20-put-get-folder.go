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

	// a key can be like this: service-name/domain
	p := &api.KVPair{Key: "service-name/domain", Value: []byte("www.example.com")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	pair, _, err := kv.Get("service-name/domain", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(pair.Key), "=>", string(pair.Value))
}
