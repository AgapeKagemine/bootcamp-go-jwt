package repository

import (
	"context"

	"gowt/internal/domain"
)

var register = `--
INSERT INTO 
	users (username, password)
VALUES
	($1, $2)
`

func (repo *UserRepositoryImpl) Register(c context.Context) error {
	user := c.Value(domain.Str("user")).(*domain.User)

	registerStmt, err := repo.db.PrepareContext(c, register)
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	row := tx.StmtContext(c, registerStmt).QueryRowContext(c, user.Username, user.Password)

	if row.Err() != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
