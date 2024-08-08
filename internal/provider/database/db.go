package database

import (
	"database/sql"
	"fmt"

	"gowt/internal/provider/database/domain"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB() (db *sql.DB, err error) {
	config := domain.NewConfig()

	connString := fmt.Sprintf("postgres://%s:%d@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.Database)

	db, err = sql.Open(config.Driver, connString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
