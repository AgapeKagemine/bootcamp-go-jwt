package domain

type DBConfig struct {
	Driver   string
	Username string
	Password uint
	Host     string
	Port     uint
	Database string
}

func NewConfig() *DBConfig {
	return &DBConfig{
		Driver:   "pgx",
		Username: "training",
		Password: 1234,
		Host:     "localhost",
		Port:     5432,
		Database: "jwt_users",
	}
}
