package route

import (
	"gin-base/src/handler"
	"github.com/gin-gonic/gin"
)

func importData(r *gin.RouterGroup) {
	var (
		g = r.Group("/import")
		h = handler.ImportData()
	)

	g.POST("/products", h.ImportProduct)
}
