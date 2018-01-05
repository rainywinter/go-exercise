package main

import (
	"context"
	"log"

	"github.com/coreos/etcd/client"
)

func GetSet(c client.Client) {
	kapi := client.NewKeysAPI(c)
	resp, err := kapi.Set(context.Background(), "/foo", "bar", nil)
	if err != nil {
		panic(err)
	}

	log.Println("set resp: ", resp)

	// get
	resp, err = kapi.Get(context.Background(), "foo", nil)
	if err != nil {
		panic(err)
	}
	log.Printf("key:%v value:%v", resp.Node.Key, resp.Node.Value)
}

func Watch(c client.Client) {
	log.Println("begin watching...")
	kapi := client.NewKeysAPI(c)
	w := kapi.Watcher("name", &client.WatcherOptions{
		Recursive: true,
	})
	resp, err := w.Next(context.Background())
	if err != nil {
		panic(err)
	}
	log.Printf("key:%v new_value:%v", resp.Node.Key, resp.Node.Value)
}

func main() {
	cfg := client.Config{
		Endpoints: []string{"http://localhost:2379"},
	}
	c, err := client.New(cfg)

	if err != nil {
		panic(err)
	}

	// get set test
	GetSet(c)

	// watch
	Watch(c)
}
