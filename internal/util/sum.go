package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// GenHash takes a file path and creates an MD5 sum of the file and returns the string or an error
func GenHash(file *os.File) (string, error) {
	var hashString string
	newHash := md5.New()
	if _, err := io.Copy(newHash, file); err != nil {
		return "", err
	}
	hashInBytes := newHash.Sum(nil)[:16]
	hashString = hex.EncodeToString(hashInBytes)
	
	return hashString, nil

}