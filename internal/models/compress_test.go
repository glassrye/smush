package models

import (
	"bytes"
	"compress/gzip"
	"io"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestCompress(t *testing.T) {
	const charMap = "abcdefghijklmnop0123456789_+"
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		filename   string
		writebytes int
		comploc    string
	}{
		{
			filename:   "testdata/test_file_1.txt",
			writebytes: 1000,
			comploc:    "testdata/test_file_1.txt.gz",
		},
	}

	for _, tt := range tests {
		data := make([]byte, tt.writebytes)
		for x := 0; x < tt.writebytes; x++ {
			data[x] = charMap[rand.Intn(len(charMap))]
		}

		f, err := os.OpenFile(tt.filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if err != nil {
			t.Fatalf("failed creating test file: %v", err)
		}
		defer os.Remove(tt.filename)

		if _, err = f.Write(data); err != nil {
			t.Fatalf("failed writing test file: %v", err)
		}

		filesize := int64(len(data))
		if err := f.Close(); err != nil {
			t.Fatalf("there was an error closing the file: %v", err)
		}

		a := Archive{
			OriginLoc:   tt.filename,
			CompressLoc: tt.comploc,
		}
		defer os.Remove(a.CompressLoc)

		if err := a.Compress(); err != nil {
			t.Fatalf("there was an error with compression: %v", err)
		}

		stat, err := os.Stat(tt.comploc)
		if err != nil {
			t.Fatalf("cannot stat compress location after compression: %v", err)
		}
		if stat.Size() >= filesize {
			t.Fatalf("compression failed. comp size: %v written bytes: %v", stat.Size(), filesize)
		}

		o, err := os.Open(a.CompressLoc)
		if err != nil {
			t.Fatalf("error opening compressed file for verifications: %v", err)
		}
		defer o.Close()

		n, err := gzip.NewReader(o)
		if err != nil {
			t.Fatalf("error creating gzip reader: %v", err)
		}
		defer n.Close()

		var buf = bytes.Buffer{}
		if _, err := io.Copy(&buf, n); err != nil {
			t.Fatalf("error copying to buffer: %v", err)
		}

		if !bytes.Equal(buf.Bytes(), data) {
			t.Fatalf("there was an error with validating uncompress. Content does not match.")
		}
	}
}
