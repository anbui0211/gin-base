package routevalidation

import (
	"gin-base/internal/response"
	gincontext "gin-base/internal/util/context"
	querymodel "gin-base/src/model/query"
	requestmodel "gin-base/src/model/request"
	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	Create() gin.HandlerFunc
	All() gin.HandlerFunc
	Update() gin.HandlerFunc
	ChangeStatus() gin.HandlerFunc
}

type userImpl struct{}

func User() UserInterface {
	return userImpl{}
}

func (u userImpl) All() gin.HandlerFunc {
	return func(c *gin.Context) {
		var q querymodel.UserAll

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

func (u userImpl) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload requestmodel.UserCreate

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

func (u userImpl) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload requestmodel.UserUpdate

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

func (u userImpl) ChangeStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload requestmodel.UserChangeStatus

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
