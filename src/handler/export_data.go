package handler

import (
	"gin-base/internal/response"
	gincontext "gin-base/internal/util/context"
	querymodel "gin-base/src/model/query"
	"gin-base/src/service"

	"github.com/gin-gonic/gin"
)

type ExportDataInterface interface {
	ExportProduct(c *gin.Context)
}

type exportDataImpl struct{}

func ExportData() ExportDataInterface {
	return &exportDataImpl{}
}

// ExportProduct ...
func (h *exportDataImpl) ExportProduct(c *gin.Context) {
	var (
		ctx   = gincontext.GetContext(c)
		s     = service.ExportData()
		query = gincontext.GetQuery(c).(querymodel.Product)
	)

	if err := s.ExportProduct(ctx, query); err != nil {
		response.R400(c, nil, err)
		return
	}

	response.R200(c, "ok", "")
}
