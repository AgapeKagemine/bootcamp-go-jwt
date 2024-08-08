package repository

import (
	"context"

	"gowt/internal/domain"
)

const findByUsername = `
SELECT 
	password
FROM
	users
WHERE 
	username = $1
`

func (repo *UserRepositoryImpl) Login(c context.Context) (string, error) {
	findByUsernameStmt, err := repo.db.PrepareContext(c, findByUsername)
	if err != nil {
		return "", err
	}

	tx, err := repo.db.BeginTx(c, nil)
	if err != nil {
		return "", err
	}

	user := c.Value(domain.Str("user")).(*domain.User)

	row := tx.StmtContext(c, findByUsernameStmt).QueryRowContext(c, user.Username)

	var password string
	err = row.Scan(
		&password,
	)
	if err != nil {
		return "", err
	}

	return password, nil
}
