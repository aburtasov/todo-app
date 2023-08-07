package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponce struct {
	Message string `json:"message"`
}

func NewErrorResponce(c *gin.Context, statusCode int, message string) *error {
	logrus.Error(message)

	c.AbortWithStatusJSON(statusCode, errorResponce{Message: message})

	return nil
}

type statusResponce struct {
	Status string `json:"status"`
}
