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
	Login(c *gin.Context)
}

type authImpl struct{}

func Auth() AuthInterface {
	return &authImpl{}
}

// Register godoc
//
//	@Summary		Register user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		requestmodel.Register	true	"registerPayload"
//	@Success		200		{object}	responsemodel.Auth
//	@Router			/auth/register [post]
func (a authImpl) Register(c *gin.Context) {
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

// Login godoc
//
//	@Summary		Register user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		requestmodel.Login	true	"loginPayload"
//	@Success		200		{object}	 responsemodel.Auth
//	@Router			/auth/login [post]
func (a authImpl) Login(c *gin.Context) {
	var (
		ctx     = gincontext.GetContext(c)
		s       = service.Auth()
		payload = gincontext.GetPayload(c).(requestmodel.Login)
	)

	res, err := s.Login(ctx, payload)
	if err != nil {
		response.R400(c, nil, err)
		return
	}

	response.R200(c, gin.H{"token": res.Token}, "")
}
