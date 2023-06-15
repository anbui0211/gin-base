package main

import (
	"gin-base/src/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())

	// Initialize
	server.Bootstrap(r)

	r.Run(":8001")
}
