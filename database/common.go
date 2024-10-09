package database

import (
	"database/sql"
	"fmt"
	u "ctrl-c/util"
)

func Open() *sql.DB {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/ctrl?sslmode=%s",
		u.Unwrap(u.Dotenv("PG_USERNAME")),
		u.Unwrap(u.Dotenv("PG_PASSWORD")),
		u.Unwrap(u.Dotenv("PG_HOST")),
		u.Unwrap(u.Dotenv("PG_PORT")),
		u.Unwrap(u.Dotenv("PG_SSL_MODE")),
	)
	db := u.Unwrap(sql.Open("postgres", connStr))

	return db
}

