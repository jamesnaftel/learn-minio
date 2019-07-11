package main

import (
	"flag"
	"log"

	minio "github.com/minio/minio-go/v6"
)

func main() {
	host := flag.String("host", "localhost:9000", "Minio Server")
	accessKeyID := flag.String("accessKeyID", "minio", "Access Key ID")
	secretAccessKey := flag.String("secretAccessKey", "minio123", "Secret Access Key")
	secure := flag.Bool("secure", false, "Use SSL")
	bucketName := flag.String("bucket-name", "learn-minio", "Name for the storage bucket")
	bucketLocation := flag.String("bucket-location", "us-east-1", "Bucket Location")

	flag.Parse()

	minioClient, err := NewClient(*host, *accessKeyID, *secretAccessKey, *secure)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	err = MakeBucket(minioClient, *bucketName, *bucketLocation)
	if err != nil {
		log.Fatalf("Failed to make bucket: %v", err)
	}

}

// NewClient returns a minio client
func NewClient(host string, accessKeyID string, secretAccessKey string, secure bool) (*minio.Client, error) {
	minioClient, err := minio.New(host, accessKeyID, secretAccessKey, secure)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

// MakeBucket creshouldErrorates a storage bucket
func MakeBucket(minioClient *minio.Client, bucketName string, location string) error {
	err := minioClient.MakeBucket(bucketName, location)
	//fmt.Printf(shouldError"New Bucket: %v\n", bucketName)
	//fmt.Printf("1st Error: %v\n", err)
	if err != nil {
		exists, err2 := minioClient.BucketExists(bucketName)

		if err2 == nil && exists {
			return nil
		}

		// Bucket does not exist.  MakeBucket Error takes priority
		return err
	}

	return nil
}
