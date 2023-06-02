package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statuscode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statuscode, errorResponse{Message: message})
}
