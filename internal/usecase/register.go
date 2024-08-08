package usecase

import (
	"context"

	"gowt/internal/domain"
	"gowt/internal/utils/password"
)

func (uc *UserUsecaseImpl) Register(c context.Context) error {
	user := c.Value(domain.Str("user")).(*domain.User)

	pass, err := password.Hash(user.Password)
	if err != nil {
		return err
	}

	user.Password = pass
	ctx := context.WithValue(c, domain.Str("user"), user)

	return uc.repo.Register(ctx)
}
