package gincontext

import (
	"context"
	"github.com/gin-gonic/gin"
)

func GetContext(c *gin.Context) context.Context {
	return c.Request.Context()
}
