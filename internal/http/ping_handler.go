package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	Message string `json:"message"`
}

func pingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, PingResponse{Message: "pong"})
	}
}
