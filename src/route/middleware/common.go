package routemiddleware

import (
	"errors"
	"gin-base/internal/constant"
	"gin-base/internal/response"
	"gin-base/internal/util"
	gincontext "gin-base/internal/util/context"
	"github.com/gin-gonic/gin"
)

func ValidID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id = c.Param("id")
		if !util.ValidID(id) {
			response.R400(c, nil, errors.New(constant.ErrInvalidID))
			return
		}

		gincontext.SetParam(c, "id", id)
		c.Next()
	}
}
