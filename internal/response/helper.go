package response

import "github.com/gin-gonic/gin"

func sendResponse(c *gin.Context, httpCode int, data interface{}, success bool, errMsg string) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(httpCode, gin.H{
		"success": success,
		"data":    data,
		"message": errMsg,
		//"code":    code,
	})
}
