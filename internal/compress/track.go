package compress

import (
	"context"
	"database/sql"
	"fmt"
)

// DB is the connection to the database exported
type DB struct {
	DB *sql.DB
}

type Tracker struct {
	Connection string  `json:"connection"`
	Table string `json:"table"`
	Archive Archive
	DB *sql.DB
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