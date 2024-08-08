package usecase

import (
	"context"
	"fmt"

	"gowt/internal/domain"
	"gowt/internal/utils/jwt"
	"gowt/internal/utils/password"
)

func (uc *UserUsecaseImpl) Login(c context.Context) (string, error) {
	user := c.Value(domain.Str("user")).(*domain.User)

	hash, err := uc.repo.Login(c)
	if err != nil || hash == "" {
		return "", err
	}

	ok := password.Check(hash, user.Password)
	if !ok {
		return "", fmt.Errorf("password not match")
	}

	token, err := jwt.CreateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
