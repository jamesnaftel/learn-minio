package main

import (
	"testing"

	minio "github.com/minio/minio-go/v6"
)

// TestMakeBucket tests the logic around MakeBucket and BucketExists
func TestMakeBucket(t *testing.T) {
	minioClient, err := setup()
	if err != nil {
		t.Fatalf("Failed to create Minio Client: %v\n", err)
	}

	tests := []struct {
		name         string
		deleteBucket bool
	}{
		{"testbucket", false}, // test creation
		{"testbucket", true},  // test exists
	}

	for _, test := range tests {
		err := MakeBucket(minioClient, test.name, "us-east-1")
		if err != nil {
			t.Errorf("Test failed for bucket name (%v): %v", test.name, err)
		}

		if test.deleteBucket {
			minioClient.RemoveBucket(test.name)
		}
	}
}

// TestMakeBucketError tests the logic around MakeBucket and BucketExists
// MakeBucket uses a strict naming path in minio while BucketExists does not.
// To ensure both paths are tested, there is a strict path error using the
// "_" and a non strict error using less than 3 characters
func TestMakeBucketError(t *testing.T) {

	minioClient, err := setup()
	if err != nil {
		t.Fatalf("Failed to create Minio Client: %v\n", err)
	}

	// Both of these should return an error
	tests := []string{"test_bucket", "1"}

	for _, bucketName := range tests {
		err := MakeBucket(minioClient, bucketName, "us-east-1")
		if err == nil {
			t.Fatalf("TestMakeBucketError failed for bucketname:  %s\n", bucketName)
		}
	}
}

func setup() (*minio.Client, error) {
	host := "localhost:9000"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	secure := false

	minioClient, err := NewClient(host, accessKeyID, secretAccessKey, secure)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
