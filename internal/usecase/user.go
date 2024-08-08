package usecase

import (
	"context"

	"gowt/internal/repository"
)

type UserUsecase interface {
	Login(context.Context) (string, error)
	Register(context.Context) error
}

type UserUsecaseImpl struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{repo: repo}
}
