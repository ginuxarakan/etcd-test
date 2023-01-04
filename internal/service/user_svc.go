package service

import (
	"context"
	"ercd-test/internal/conf"
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

	telegram *TelegramBot
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
	opts := []grpc.DialOption{
		grpc.WithResolvers(etcdResolver),
		grpc.WithDefaultServiceConfig(svcConfig),
		grpc.WithInsecure(),
		//grpc.WithBlock(),
	}
	logger.Logrus.Info(etcdResolver)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	key := fmt.Sprintf("%s:///%s", conf.Etcd().Scheme, conf.RPCSvc().UserRPC.Name)
	conn, err := grpc.DialContext(ctx, key, opts...)
	if err != nil {
		logger.Logrus.Error(err)
		return nil, err
	}

	client := pb.NewUserServiceClient(conn)

	logger.Logrus.Println("Connected to User RPC")

	return &UserSvc{
		Conn:     conn,
		Client:   client,
		telegram: NewTelegramBot(conf.Telegram().TokenID, conf.Telegram().GroupID),
	}, nil
}

func (s *UserSvc) Run() {
	s.listenBalanceChange()
}

func (s *UserSvc) listenBalanceChange() {
	stream, err := s.Client.StreamTest(context.Background(), &pb.StreamTestReq{})
	if err != nil {
		logger.Logrus.Error(err)
		return
	}

	go func() {
		defer logger.Logrus.Info("Streaming end ...")
		for {
			value, err := stream.Recv()
			if err != nil {
				logger.Logrus.Error(err)
				return
			}

			fmt.Println(value)
			go func() {
				fmt.Println("send to telegram")

				msg := fmt.Sprintf("Stream: %v", value.Message)
				_, err := s.telegram.SendMessage(msg)
				if err != nil {
					logger.Logrus.Error(err)
					return
				}
				logger.Logrus.Debug("send to telegram success")
			}()

		}
	}()
}

func (s *UserSvc) Close() {
	s.Conn.Close()
}
