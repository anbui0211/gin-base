package main

import (
	_ "gin-base/docs"
	"gin-base/src/server"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           			Golang Project
// @version         			1.0

// @contact.name    			Bui Phu An
// @contact.github 				https://github.com/buiphuan0211
// @contact.linked 				https://www.linkedin.com/in/buiphuan/
// @host      					localhost:8001
// @BasePath  					/api/v1
// @securityDefinitions.basic  	BasicAuth
func main() {
	r := gin.Default()

	r.Use(gin.Recovery())

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Some errorcode occured. Err: %s", err)
	}

	// Initialize
	server.Bootstrap(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(os.Getenv("APP_PORT"))
}
