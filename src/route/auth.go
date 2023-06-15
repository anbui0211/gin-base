package route

import (
	"gin-base/src/handler"
	routevalidation "gin-base/src/route/validation"
	"github.com/gin-gonic/gin"
)

func auth(r *gin.RouterGroup) {
	var (
		g = r.Group("/auth")
		h = handler.Auth()
		v = routevalidation.Auth()
	)

	g.POST("/register", v.Register(), h.Register)
	//g.POST("/login", v.Register(), h.Login)

}
