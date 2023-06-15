package routevalidation

import (
	"gin-base/internal/response"
	gincontext "gin-base/internal/util/context"
	requestmodel "gin-base/src/model/request"
	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Register() gin.HandlerFunc
}
type authImpl struct{}

func Auth() AuthInterface {
	return authImpl{}
}

func (a authImpl) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload requestmodel.Register
		if err := c.ShouldBindJSON(&payload); err != nil {
			response.R400(c, nil, err)
			return
		}

		if err := payload.Validate(); err != nil {
			response.R400(c, nil, err)
			return
		}

		gincontext.SetPayload(c, payload)
		c.Next()
	}
}
