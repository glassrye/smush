package models

import (
	"bytes"
	"compress/gzip"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/deepdyve/compress-logs/internal/util"
)

// DB is the connection to the database exported
var DB *sql.DB

// Define a archive type and the attributes it can have

type Archive struct {
	Archived bool `json:"archived"`
	Online bool	 `json:"online"`
	OriginHost string `json:"origin_host"`
	OriginLoc string `json:"disk_loc,omitempty"`
	CompressLoc string `json:"compress_loc,omitempty"`
	ArchiveBucket string `json:"archive_bucket"`
	ArchivePrefix string `json:"archive_prefix"`
	OriginHash string `json:"origin_hash"`
	CompressHash string `json:"compress_hash"`
	LastUpdate string   `json:"last_update"`
}


func AddRecord(o, c, on, cn string) (sql.Result, error) {

	stmt := "INSERT INTO `filestat` (`orig_sum`, `comp_sum`, `oname`, `cname`, `cur_loc`) VALUES(?, ?, ?, ?, ?)"
	res, err := DB.ExecContext(context.Background(), stmt, o, c, on, cn, cn)
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


// Compress is a receiver of type Archive
// it expect a pointer as we add the the value of OriginHash and CompressHash
// these are tracked variables in the DB
func (a *Archive) Compress() error {
	fmt.Printf("Orig: %v  Comp: %v\n", a.OriginLoc, a.CompressLoc) 

    comp, err := os.OpenFile(a.CompressLoc, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(a.OriginLoc)
	if err != nil {
		return err
	}

	oFile, err := os.Open(a.OriginLoc)
	if err != nil {
		return nil
	}
	oHash, err := util.GenHash(oFile)
	if err != nil {
		return err
	}
	if err = oFile.Close(); err != nil {
		return err
	}

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	zw.Write(data)
	zw.Close()

	if _, err := io.Copy(comp, &buf); err != nil {
		fmt.Printf("error copying data from buffer: %v", err)
	}
	if err = os.Remove(a.OriginLoc); err != nil {
		fmt.Printf("there was an error removing the file: %s\n", a.OriginLoc)
		return err
	}
	gFile, err := os.Open(a.CompressLoc)
	if err != nil {
		return err
	}
	gHash, err := util.GenHash(gFile)
	if err != nil {
		return err
	}
	a.CompressHash = gHash
	a.OriginHash = oHash

	return nil
}