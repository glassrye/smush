package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// NewDatabase is used when a Tracker is created typically. However, it just returns a pointer
// to sql.DB and an error based on username, password, database name, hostname string inputs
func NewDatabase(u, p, dbn, h string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=5432", u, p, dbn, h)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
