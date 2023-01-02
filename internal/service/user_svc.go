package service

import (
	"context"
	"ercd-test/internal/logger"
	"ercd-test/internal/pb"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"time"
)

type UserSvc struct {
	Conn   *grpc.ClientConn
	Client pb.UserServiceClient
}

func NewUserService() (*UserSvc, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		logger.Logrus.Error(err)
		return nil, err
	}

	etcdResolver, err := resolver.NewBuilder(cli)
	if err != nil {
		logger.Logrus.Error(err)
		return nil, err
	}

	svcConfig := fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)

	logger.Logrus.Info(svcConfig)

	opts := []grpc.DialOption{
		grpc.WithResolvers(etcdResolver),
		grpc.WithDefaultServiceConfig(svcConfig),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}
	logger.Logrus.Info(etcdResolver)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "etcd:///userRPC", opts...)
	if err != nil {
		logger.Logrus.Error(err)
		return nil, err
	}

	client := pb.NewUserServiceClient(conn)

	logger.Logrus.Println("Connected to User RPC")

	return &UserSvc{
		Conn:   conn,
		Client: client,
	}, nil
}

func (s *UserSvc) Close() {
	s.Conn.Close()
}
