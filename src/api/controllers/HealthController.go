package controllers

import (
	"github.com/gin-gonic/gin"
)

type HealthController interface {
	Ping(c *gin.Context)
}

type healtController struct{}

func NewHealthController() HealthController {
	return &healtController{}
}

func (hc *healtController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
