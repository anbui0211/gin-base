package route

import (
	"gin-base/src/handler"
	routemiddleware "gin-base/src/route/middleware"
	routevalidation "gin-base/src/route/validation"

	"github.com/gin-gonic/gin"
)

func user(r *gin.RouterGroup) {
	var (
		g = r.Group("users")
		h = handler.User()
		v = routevalidation.User()
	)

	// Create
	g.POST("", v.Create(), h.Create)

	// All
	g.GET("", v.All(), h.All)

	// Detail
	g.GET("/:id", routemiddleware.ValidID(), h.Detail)

	// Update
	g.PUT("/:id", routemiddleware.ValidID(), v.Update(), h.Update)

	// Update status
	g.PATCH("/:id/status", routemiddleware.ValidID(), v.ChangeStatus(), h.ChangeStatus)

}
