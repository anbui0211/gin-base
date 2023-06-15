package handler

import (
	"gin-base/internal/response"
	gincontext "gin-base/internal/util/context"
	requestmodel "gin-base/src/model/request"
	"gin-base/src/service"

	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Register(c *gin.Context)
}

type AuthImpl struct{}

func Auth() AuthInterface {
	return AuthImpl{}
}

func (a AuthImpl) Register(c *gin.Context) {
	var (
		ctx     = gincontext.GetContext(c)
		s       = service.Auth()
		payload = gincontext.GetPayload(c).(requestmodel.Register)
	)

	res, err := s.Register(ctx, payload)
	if err != nil {
		response.R400(c, nil, err)
		return
	}

	response.R200(c, gin.H{"token": res.Token}, "")
}
