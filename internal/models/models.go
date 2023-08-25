package models

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB is the connection to the database exported
type DB struct {
	DB *sql.DB
}

// Define a archive type and the attributes it can have

type Archive struct {
	Archived      bool   `json:"archived"`
	Online        bool   `json:"online"`
	OriginHost    string `json:"origin_host"`
	OriginLoc     string `json:"disk_loc,omitempty"`
	CompressLoc   string `json:"compress_loc,omitempty"`
	ArchiveBucket string `json:"archive_bucket"`
	ArchivePrefix string `json:"archive_prefix"`
	OriginHash    string `json:"origin_hash"`
	CompressHash  string `json:"compress_hash"`
	LastUpdate    string `json:"last_update"`
}

type Tracker struct {
	Connection string  `json:"connection"`
	Table string `json:"table"`
	Archive Archive
	DB *sql.DB
}

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

// AddRecord - Is a method put on an archive type that puts the results into the database
func (t *Tracker) AddRecord() (sql.Result, error) {
	stmt := `INSERT INTO filestat (orig_sum, comp_sum, oname, cname, cur_loc) VALUES ($1, $2, $3, $4, $5)`
	res, err := t.DB.ExecContext(context.Background(), stmt, t.Archive.OriginHash, t.Archive.CompressHash, t.Archive.OriginLoc, t.Archive.CompressLoc, t.Archive.OriginLoc)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("there was a weird error that I'm not sure is fatal: %v", err)
	}
	fmt.Printf("Last Insert ID: %v\n", id)

	return res, nil
}
