package server

import (
	"context"
	"ercd-test/interanl/logger"
	"ercd-test/interanl/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) UserCallTest(ctx context.Context, in *pb.UserReq) (*pb.UserResp, error) {
	logger.Logrus.Debug("From User RPC")
	return &pb.UserResp{}, nil
}
