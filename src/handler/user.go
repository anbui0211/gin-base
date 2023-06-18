package handler

import (
	"gin-base/internal/response"
	gincontext "gin-base/internal/util/context"
	querymodel "gin-base/src/model/query"
	requestmodel "gin-base/src/model/request"
	responsemodel "gin-base/src/model/response"
	"gin-base/src/service"

	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	Create(c *gin.Context)
	All(c *gin.Context)
	Detail(c *gin.Context)
	Update(c *gin.Context)
	ChangeStatus(c *gin.Context)
}

type userImpl struct{}

func User() UserInterface {
	return userImpl{}
}

// Create godoc
//
//	@Summary		Create user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		requestmodel.UserCreate	true	"userPayload"
//	@Success		200		{object}	 responsemodel.Upsert
//	@Router			/users [post]
func (userImpl) Create(c *gin.Context) {
	var (
		ctx           = gincontext.GetContext(c)
		payloadCreate = gincontext.GetPayload(c).(requestmodel.UserCreate)
		s             = service.User()
	)

	res, err := s.Create(ctx, payloadCreate)
	if err != nil {
		response.R400(c, nil, err)
		return
	}

	response.R200(c, gin.H{"id": res.ID}, "")
	return
}

// All godoc
// @Summary      List user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param 		 payload query querymodel.UserAll true "Query"
// @Success      200  {object}   responsemodel.UserAll
// @Router       /users [get]
func (h userImpl) All(c *gin.Context) {
	var (
		ctx   = gincontext.GetContext(c)
		query = gincontext.GetQuery(c).(querymodel.UserAll)
		s     = service.User()
	)

	res := s.All(ctx, query)
	response.R200(c, gin.H{"data": res}, "")
	return
}

// Detail godoc
// @Summary      Detail user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param  id path string true "userId"
// @Success      200  {object}   responsemodel.UserDetail
// @Router       /users/{id} [get]
func (userImpl) Detail(c *gin.Context) {
	var (
		id  = gincontext.GetParam(c, "id").(string)
		s   = service.User()
		ctx = gincontext.GetContext(c)
	)

	res, err := s.Detail(ctx, id)
	if err != nil {
		response.R400(c, nil, err)
		return
	}
	response.R200(c, gin.H{"data": res}, "")
}

// Update godoc
//
//	@Summary		Update user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//
// @Param  id path string true "userId"
//
//	@Param			payload	body		requestmodel.UserUpdate	true	"userPayload"
//	@Success		200		{object}	 responsemodel.Upsert
//	@Router			/users/{id} [put]
func (userImpl) Update(c *gin.Context) {
	var (
		id      = gincontext.GetParam(c, "id").(string)
		payload = gincontext.GetPayload(c).(requestmodel.UserUpdate)
		s       = service.User()
		ctx     = gincontext.GetContext(c)
	)

	res, err := s.Update(ctx, id, payload)
	if err != nil {
		response.R400(c, nil, err)
		return
	}

	response.R200(c, responsemodel.Upsert{ID: res.ID}, "")
	return
}

// ChangeStatus godoc
//
//	@Summary		Change Status User
//	@Tags			users
//	@Accept			json
//	@Produce		json
//
// @Param  id path string true "userId"
//
//	@Param			payload	body		requestmodel.UserChangeStatus	true	"userPayload"
//	@Success		200		{object}	 responsemodel.Upsert
//	@Router			/users/{id}/status [patch]
func (userImpl) ChangeStatus(c *gin.Context) {
	var (
		id      = gincontext.GetParam(c, "id").(string)
		payload = gincontext.GetPayload(c).(requestmodel.UserChangeStatus)
		s       = service.User()
		ctx     = gincontext.GetContext(c)
	)

	res, err := s.ChangeStatus(ctx, id, payload)
	if err != nil {
		response.R400(c, nil, err)
		return
	}

	response.R200(c, responsemodel.Upsert{ID: res.ID}, "")
	return
}
