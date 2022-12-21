package main

import (
	"context"
	"fmt"
	client "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	dailTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	cli, _ := client.New(client.Config{
		DialTimeout: dailTimeout,
		Endpoints:   []string{"127.0.0.1:2379"},
	})
	defer cli.Close()

	kv := client.NewKV(cli)

	GetSingleValueDemo(ctx, kv)
}

func GetSingleValueDemo(ctx context.Context, kv client.KV) {
	fmt.Println("Get single Value Demo ........... ")

	// Delete all keys
	kv.Delete(ctx, "micky", client.WithPrefix())

	// Insert a key value
	pr, _ := kv.Put(ctx, "micky", "Diamond 777")
	rev := pr.Header.Revision
	fmt.Println("Revision: ", rev)

	gr, _ := kv.Get(ctx, "micky")
	fmt.Println("GR Values : ", gr.Kvs)
	fmt.Println("Value: ", string(gr.Kvs[0].Value), " Revision: ", gr.Header.Revision)

	// modify the value of an existing key (create new reversion)
	kv.Put(ctx, "micky", "Diamond 900")

	gr, _ = kv.Get(ctx, "micky")
	fmt.Println("Value: ", string(gr.Kvs[0].Value), " Revision: ", gr.Header.Revision)

	// Get the value of the previous revision
	gr, _ = kv.Get(ctx, "micky", client.WithRev(rev))
}
