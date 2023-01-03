package server

import (
	"context"
	"ercd-test/internal/dto"
	"ercd-test/internal/logger"
	"ercd-test/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	messageChan chan *dto.Message

	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) UserCallTest(ctx context.Context, in *pb.UserReq) (*pb.UserResp, error) {
	logger.Logrus.Debug("From User RPC")
	return &pb.UserResp{}, nil
}

func (s *UserService) StreamInput(ctx context.Context, in *pb.StreamInputReq) (*pb.StreamInputResp, error) {
	message := &dto.Message{}
	if in.Input != "" {
		message.Message = in.Input
	}

	s.messageChan <- message

	return &pb.StreamInputResp{}, nil
}

func (s *UserService) StreamTest(in *pb.StreamTestReq, stream pb.UserService_StreamTestServer) error {
	defer logger.Logrus.Info("Stream ended ... ")
	logger.Logrus.Info("Started Message subscribe ... ")

	for {
		select {
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream has ended.")
		case m := <-s.messageChan:
			err := stream.SendMsg(&pb.StreamTestResp{
				Message: m.Message,
			})
			if err != nil {
				logger.Logrus.Error(err)
				return err
			}
		}
	}
}
