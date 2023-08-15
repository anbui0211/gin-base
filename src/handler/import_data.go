package handler

import (
	"gin-base/internal/response"
	gincontext "gin-base/internal/util/context"
	"gin-base/src/service"

	"github.com/gin-gonic/gin"
)

type ImportDataInterface interface {
	ImportProduct(c *gin.Context)
}

type importDataImpl struct{}

func ImportData() ImportDataInterface {
	return &importDataImpl{}
}

func (i importDataImpl) ImportProduct(c *gin.Context) {
	var (
		ctx      = gincontext.GetContext(c)
		s        = service.ImportData()
		filename = c.Query("filename")
	)

	data, err := s.ImportProducts(ctx, filename)
	if err != nil {
		response.R400(c, nil, err)
		return
	}

	response.R200(c, gin.H{"products": data}, "")
}
