package models

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/deepdyve/compress-logs/internal/util"
)

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