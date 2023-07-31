package routevalidation

import (
	"gin-base/internal/response"
	gincontext "gin-base/internal/util/context"
	querymodel "gin-base/src/model/query"

	"github.com/gin-gonic/gin"
)

type ExportDataInterface interface {
	ExportProduct() gin.HandlerFunc
}

type exportDataImpl struct{}

func ExportData() ExportDataInterface {
	return &exportDataImpl{}
}

func (v *exportDataImpl) ExportProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var q querymodel.Product

		if err := c.ShouldBind(&q); err != nil {
			response.R400(c, nil, err)
			return
		}

		if err := q.Validate(); err != nil {
			response.R400(c, nil, err)
			return
		}

		gincontext.SetQuery(c, q)
		c.Next()
	}
}
