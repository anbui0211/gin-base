package route

import (
	"gin-base/src/handler"
	routevalidation "gin-base/src/route/validation"

	"github.com/gin-gonic/gin"
)

func exportData(r *gin.RouterGroup) {
	var (
		// Export pdf file
		g = r.Group("/export-pdf")
		h = handler.ExportData()
		v = routevalidation.ExportData()
	)

	g.GET("/products", v.ExportProduct(), h.ExportProduct)
}
