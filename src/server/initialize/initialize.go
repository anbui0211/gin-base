package initialize

import "github.com/gin-gonic/gin"

func Init(g *gin.Engine) {
	mySQL()
}
