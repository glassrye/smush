package models

import (
	"context"
	"database/sql"
	"log"
)

// DB is the connection to the database exported
var DB *sql.DB

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

// AddRecord - Is a method put on an archive type that puts the results into the database
func (a *Archive) AddRecord() (sql.Result, error) {
	stmt := "INSERT INTO `filestat` (`orig_sum`, `comp_sum`, `oname`, `cname`, `cur_loc`) VALUES(?, ?, ?, ?, ?)"
	res, err := DB.ExecContext(context.Background(), stmt, a.OriginHash, a.CompressHash, a.OriginLoc, a.CompressLoc, a.OriginLoc)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows != 1 {
		log.Fatalf("expected to insert 1 row, but got result count of %d", rows)
	}

	return res, nil
}
