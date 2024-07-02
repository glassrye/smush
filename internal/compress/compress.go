package compress

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/deepdyve/compress-logs/internal/util"
)
type Archive struct {
	Archive     bool   `json:"archived"`
	Online        bool   `json:"online"`
	Recurse       bool   `json:"recurse"`
	OriginHost    string `json:"origin_host"`
	OriginLoc     string `json:"disk_loc,omitempty"`
	CompressLoc   string `json:"compress_loc,omitempty"`
	ArchiveBucket string `json:"archive_bucket"`
	ArchivePrefix string `json:"archive_prefix"`
	OriginHash    string `json:"origin_hash"`
	CompressHash  string `json:"compress_hash"`
	LastUpdate    string `json:"last_update"`
	DBHost string `json:"db_host,omitempty"`
	DBUser string `json:"db_user,omitempty"`
	DBPass string `json:"_,omitempty"`
	Match string `json:"match,omitempty"`
	Suffix string `json:"suffix,omitempty"`
	Providers []string `json:"providers,omitempty"`
	Level int `json:"level,omitempty"`
}

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
	hashFile, err := os.Open(a.CompressLoc)
	if err != nil {
		fmt.Printf("error opening compress file: %v", err)
		return err
	}
	defer hashFile.Close()
	dstHash, err := util.GenHash(hashFile)
	if err != nil {
		fmt.Printf("error generating hash for %s with error: %v", a.CompressLoc, hashFile.Name())
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
