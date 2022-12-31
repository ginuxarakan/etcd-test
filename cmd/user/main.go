package main

import (
	"context"
	"ercd-test/cmd/user/server"
	"ercd-test/interanl/logger"
	"ercd-test/interanl/pb"
	"flag"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := flag.String("port", "8070", "UserRPC Default port is 8070.")
	flag.Parse()

	addr := fmt.Sprintf("0.0.0.0:%s", *port)

	//  grpc start
	gs := grpc.NewServer()

	// create user service instance
	userSvc := server.NewUserService()
	pb.RegisterUserServiceServer(gs, userSvc)

	// listening
	l, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Logrus.Fatalf("Failed to Listen User RPC server: %v", err)
	}
	defer l.Close()

	// service register with etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		logger.Logrus.Fatalf("Failed to connet with etcd: %v", err)
	}

	em, err := endpoints.NewManager(cli, "userRPC")
	if err != nil {
		logger.Logrus.Panic(err)
	}

	if err := em.AddEndpoint(context.Background(), "userRPC/"+addr, endpoints.Endpoint{
		Addr: addr,
	}); err != nil {
		logger.Logrus.Panic(err)
	}

	// start grpc server
	go func() {
		logger.Logrus.Println("User RPC server is starting on: ", *port)
		if err := gs.Serve(l); err != nil {
			logger.Logrus.Fatalf("Failed to start server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	cli.Close()
	gs.Stop()
}
