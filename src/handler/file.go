package handler

import (
	"gin-base/internal/response"
	gincontext "gin-base/internal/util/context"
	"gin-base/src/service"
	"github.com/gin-gonic/gin"
)

// FileDataInterface ...
type FileDataInterface interface {
	UploadProducts(c *gin.Context)
}

// fileDataImpl ...
type fileDataImpl struct{}

// FileData ...
func FileData() FileDataInterface {
	return &fileDataImpl{}
}

// UploadProducts ...
func (v *fileDataImpl) UploadProducts(c *gin.Context) {
	var (
		s   = service.FileData()
		ctx = gincontext.GetContext(c)
	)

	err := s.UploadProducts(ctx)
	if err != nil {
		response.R400(c, nil, err)
		return
	}

	response.R200(c, nil, "oke")
}
