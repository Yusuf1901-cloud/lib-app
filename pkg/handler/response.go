package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, messase string) {
	logrus.Error(messase)

	c.AbortWithStatusJSON(statusCode, errorResponse{Message: messase})
}
