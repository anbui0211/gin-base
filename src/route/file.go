package route

import (
	"gin-base/src/handler"
	routevalidation "gin-base/src/route/validation"
	"github.com/gin-gonic/gin"
)

func fileData(r *gin.RouterGroup) {
	var (
		gUp = r.Group("/upload")
		v   = routevalidation.FileData()
		h   = handler.FileData()
	)

	gUp.GET("/products", v.Upload(), h.UploadProducts)
}
