package repository

import (
	"context"
	"database/sql"
)

type UserRepository interface {
	Login(context.Context) (string, error)
	Register(context.Context) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}
