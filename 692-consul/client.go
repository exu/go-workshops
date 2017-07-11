package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

// main  docker run -d --name=dev-consul consul
func main() {
	// Get a new client
	config := api.DefaultConfig()
	config.Address = "172.17.0.2:8500"
	fmt.Printf("%+v", config)
	client, err := api.NewClient(config)
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
	fmt.Printf("KV: %v", pair)
}
