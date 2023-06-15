package gincontext

import "github.com/gin-gonic/gin"

func GetPayload(c *gin.Context) interface{} {
	return c.MustGet("payload")
}

func SetPayload(c *gin.Context, value interface{}) {
	c.Set("payload", value)
}

func GetQuery(c *gin.Context) interface{} {
	return c.MustGet("query")
}

func SetQuery(c *gin.Context, value interface{}) {
	c.Set("query", value)
}

func SetParam(c *gin.Context, key string, value interface{}) {
	c.Set(key, value)
}

func GetParam(c *gin.Context, key string) interface{} {
	return c.MustGet(key)
}
