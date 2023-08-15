package handler

import (
	"github.com/gin-gonic/gin"
)

type Common struct{}

func (Common) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
