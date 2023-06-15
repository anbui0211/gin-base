package gincontext

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func GetContext(c *gin.Context) context.Context {
	return c.Request.Context()
}
