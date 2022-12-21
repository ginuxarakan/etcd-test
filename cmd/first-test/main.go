package main

import (
	"context"
	"fmt"
	client "go.etcd.io/etcd/client/v3"
	"strconv"
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

	//GetSingleValueDemo(ctx, kv)

	//GetMultipleValuesWithPaginationDemo(ctx, kv)

	LeaseDemo(ctx, cli, kv)
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

func GetMultipleValuesWithPaginationDemo(ctx context.Context, kv client.KV) {
	fmt.Println("Get multiple values demo .............. ")

	// Delete all keys
	kv.Delete(ctx, "test:key:", client.WithPrefix())

	// Insert multiple keys
	for i := 0; i < 20; i++ {
		k := fmt.Sprintf("test:key:%02d", i)
		kv.Put(ctx, k, strconv.Itoa(i))
	}

	opts := []client.OpOption{
		client.WithFromKey(),
		client.WithSort(client.SortByKey, client.SortAscend),
		client.WithLimit(5),
	}

	gr, _ := kv.Get(ctx, "test:key:", opts...)
	fmt.Println("First Page ---------------- ")
	for _, item := range gr.Kvs {
		fmt.Println(string(item.Key), " ", string(item.Value))
	}

	lastKey := string(gr.Kvs[len(gr.Kvs)-1].Key)

	fmt.Println("last key: ", lastKey)

	fmt.Println("Second Page --------------- ")
	opts = append(opts, client.WithFromKey())
	gr, _ = kv.Get(ctx, lastKey, opts...)

	// skipping the first item, which the last item from the previous Get
	for _, item := range gr.Kvs[1:] {
		fmt.Println(string(item.Key), " ", string(item.Value))
	}
}

func GetMultipleValuesWithPaginationDemoExp(ctx context.Context, kv client.KV) {
	fmt.Println("*** GetMultipleValuesWithPaginationDemo()")
	// Delete all keys
	kv.Delete(ctx, "key", client.WithPrefix())

	// Insert 20 keys
	for i := 0; i < 20; i++ {
		k := fmt.Sprintf("key_%02d", i)
		kv.Put(ctx, k, strconv.Itoa(i))
	}

	opts := []client.OpOption{
		client.WithFromKey(),
		client.WithSort(client.SortByKey, client.SortAscend),
		client.WithLimit(3),
	}

	gr, _ := kv.Get(ctx, "key", opts...)

	fmt.Println("--- First page ---")
	for _, item := range gr.Kvs {
		fmt.Println(string(item.Key), string(item.Value))
	}

	lastKey := string(gr.Kvs[len(gr.Kvs)-1].Key)

	fmt.Println("--- Second page ---")
	opts = append(opts, client.WithFromKey())
	gr, _ = kv.Get(ctx, lastKey, opts...)

	// Skipping the first item, which the last item from from the previous Get
	for _, item := range gr.Kvs[1:] {
		fmt.Println(string(item.Key), string(item.Value))
	}
}

func LeaseDemo(ctx context.Context, cli *client.Client, kv client.KV) {
	fmt.Println("Lease Demo -------------- ")

	// Delete all keys
	kv.Delete(ctx, "test:micky", client.WithPrefix())

	gr, _ := kv.Get(ctx, "test:micky")
	if len(gr.Kvs) == 0 {
		fmt.Println("There's no keys.")
	}

	lease, _ := cli.Grant(ctx, 1)

	// Insert key with a lease of 1 second TTL
	kv.Put(ctx, "test:micky", "diamond micky 777", client.WithLease(lease.ID))

	gr, _ = kv.Get(ctx, "test:micky")
	if len(gr.Kvs) == 1 {
		fmt.Println("Found Key.")
	}

	// Let the TTL expire
	time.Sleep(time.Second * 3)

	gr, _ = kv.Get(ctx, "test:micky")
	if len(gr.Kvs) == 0 {
		fmt.Println("No more keys.")
	}
}
