package route

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	common(v1)
	user(v1)
	auth(v1)
}
