package routevalidation

import (
	"github.com/gin-gonic/gin"
)

type FileDataInterface interface {
	Upload() gin.HandlerFunc
}

type fileDataImpl struct{}

func FileData() FileDataInterface {
	return &fileDataImpl{}
}

func (v *fileDataImpl) Upload() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
