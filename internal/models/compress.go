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
// it expect a pointer as we add the value of OriginHash and CompressHash
// these can be tracked as variables in the DB
func (a *Archive) Compress() error {
	// Read the data from the original fine into a byte slice
	srcData, err := os.ReadFile(a.OriginLoc)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(a.OriginLoc)
	if err != nil {
		return err
	}

	dstFile, err := os.OpenFile(a.CompressLoc, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	// Create a hash from that original file
	srcHash, err := util.GenHash(srcFile)
	if err != nil {
		return err
	}

	if err = srcFile.Close(); err != nil {
		return err
	}

	var buf bytes.Buffer

	gzipWriter := gzip.NewWriter(&buf)
	gzipWriter.Write(srcData)

	if err = gzipWriter.Flush(); err != nil {
		return err
	}

	if err = gzipWriter.Close(); err != nil {
		return err
	}

	if _, err := io.Copy(dstFile, &buf); err != nil {
		fmt.Printf("error writing file: %v", err)
		return err
	}

	dstHash, err := util.GenHash(dstFile)
	if err != nil {
		return err
	}

	if err = dstFile.Close(); err != nil {
		fmt.Println("erorr closing file: ", err)
		return err
	}
	// Set the hashes in to the archive struct and return a nil error
	a.CompressHash = dstHash
	a.OriginHash = srcHash

	fmt.Println("Original Hash ", a.OriginHash, "Compress hash ", a.CompressHash)
	if err = os.Remove(a.OriginLoc); err != nil {
		fmt.Println("error removing original file: ", err)
		return err
	}
	return nil
}
