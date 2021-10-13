package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// NewDB constructs a Trades value for managing stock trades in a
// SQLite database. This API is not thread safe.
func NewDB(dbFile string) (*sqlx.DB, error) {
	sqlDB, err := sqlx.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	return sqlDB, nil
}
