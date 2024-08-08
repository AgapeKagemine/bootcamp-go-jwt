package user

import (
	"gowt/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Login(*gin.Context)
	Register(*gin.Context)
}

type UserHandlerImpl struct {
	uc usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) UserHandler {
	return &UserHandlerImpl{
		uc: uc,
	}
}
