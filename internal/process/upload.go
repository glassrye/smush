package process

import "fmt"

// Upload - takes a file path, a bucket name, and a bucket path as strings and returns an error
// this utility uploads  the file at 'file' to bucket at 'bucket' with bucket path 'path'
func Upload(file, bucket, prefix string) error {

	fmt.Printf("Uploading %s, to %s with path %s\n", file, bucket, prefix)
	us := fmt.Sprintf("gs://%s/%s", bucket, prefix)
	fmt.Printf("Final Upload String: %s\n", us)
	return nil
}