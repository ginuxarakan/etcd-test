package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	R *gin.Engine
}

type HConfig struct {
	R *gin.Engine
}

func NewHandler(c *HConfig) (*Handler, error) {
	return &Handler{
		R: c.R,
	}, nil
}

func (h *Handler) Register() {
	h.R.GET(
		"/",
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": "Success",
			})
		},
	)
}
