package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"megavault/api/config"
)

func NewPGSQLStorage() (*sql.DB, error) {
	conninfo := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.Env.DBUser,
		config.Env.DBPassword,
		config.Env.DBHost,
		config.Env.DBPort,
		config.Env.DBName,
	)

	print(conninfo, "\n")

	return sql.Open("postgres", conninfo)
}
