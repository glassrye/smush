package process

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/deepdyve/compress-logs/internal/util"
)

// Compress - takes a string that is a path to a file. it then compresses the file and returns
// the gzip file path, the original file hash, the gzip file hash, and an error

// With a refactor but that has to be in the models package. You can't define new methods outside
// the original package that exports the type
// compress could take a pointer to an Archive type and only return an error
// then I can add in a channel that takes an error only. We can assume it worked
// and we can add a logger to the Archive so we don't have to assume
func Compress(f string) (string, string, string, error) {
	// Create gzip filepath string and "open" it
	comPath := f + ".gz"
	comp, err := os.OpenFile(comPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return "", "", "", fmt.Errorf("error creating comp path file: %v", err)
	}

	// Read the file data into a byte slice
	data, err := os.ReadFile(f)
	if err != nil {
		return "", "", "", err
	}
	fmt.Printf("Read %d bytes from %s\n", len(data), f)

	oFile, err := os.Open(f)
	if err != nil {
		return "", "", "", err
	}
	oHash, err := util.GenHash(oFile)
	if err != nil {
		return "", "", "", err
	}
	if err = oFile.Close(); err != nil {
		return "", "", "", err
	}
	// Create a bytes buffer and set it for the new gzip writer
	// Write the data from ReadFile into the buffer and then close
	// the writer
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	zw.Write(data)
	zw.Close()

	// Now copy the buffer data into the gzip file created at the top
	// of this function
	if _, err := io.Copy(comp, &buf); err != nil {
		fmt.Printf("error copying data from buffer: %v", err)
	}
	if err = os.Remove(f); err != nil {
		fmt.Printf("there was a problem removing file: %s\n", f)
		return "", "", "", err
	}

	gFile, err := os.Open(comPath)
	if err != nil {
		return "", "", "", err
	}
	gHash, err := util.GenHash(gFile)
	if err != nil {
		return "", "", "", err
	}
	// Return the new gziped file path
	return comPath, oHash, gHash, nil
}
