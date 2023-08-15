package route

import (
	middlewareinternal "gin-base/internal/util/middlewares"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	// Middleware ...
	r.Use(middlewareinternal.CORSMiddleware())

	v1 := r.Group("/api/v1")

	common(v1)
	user(v1)
	auth(v1)
	importData(v1)
}
