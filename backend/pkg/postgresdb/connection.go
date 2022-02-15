package postgresdb

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib" // driver for connection.
	"github.com/jmoiron/sqlx"
)

// NewClient connect to postgres and ping it. Return connection to database.
func NewClient(host, port, user, password, dbName string) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := sqlx.Connect("pgx", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
