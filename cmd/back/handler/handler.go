package handler

import (
	"context"
	"ercd-test/internal/logger"
	"ercd-test/internal/pb"
	"ercd-test/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	R       *gin.Engine
	UserSvc *service.UserSvc
}

type HConfig struct {
	R *gin.Engine
}

func NewHandler(c *HConfig) (*Handler, error) {
	userSvc, err := service.NewUserService()
	if err != nil {
		logger.Logrus.Error(err)
		return nil, err
	}

	return &Handler{
		R:       c.R,
		UserSvc: userSvc,
	}, nil
}

func (h *Handler) Register() {
	h.R.GET(
		"/",
		func(c *gin.Context) {

			if _, err := h.UserSvc.Client.UserCallTest(context.Background(), &pb.UserReq{}); err != nil {
				c.JSON(500, gin.H{
					"data": err.Error(),
				})
			}

			c.JSON(200, gin.H{
				"data": "Success",
			})
		},
	)
}
