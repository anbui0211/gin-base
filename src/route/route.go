package route

import (
	routemiddleware "gin-base/src/route/middleware"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	// Middleware ...
	r.Use(routemiddleware.CORSMiddleware())

	v1 := r.Group("/api/v1")

	common(v1)
	user(v1)
	auth(v1)
}
