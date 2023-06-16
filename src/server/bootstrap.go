package server

import (
	"gin-base/src/route"
	"gin-base/src/server/initialize"
	"github.com/gin-gonic/gin"
	"log"
)

// Bootstrap ...
func Bootstrap(r *gin.Engine) {
	log.Println("Initializing server ...")
	initialize.Init(r)
	route.Init(r)
}
