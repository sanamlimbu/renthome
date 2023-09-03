package renthome

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/ninja-software/terror/v2"
)

// ObjectStorage contains the minio client
type ObjectStorage struct {
	*minio.Client
	BucketName     string
	BucketLocation string
}

// NewObjectStorage will set up the minioClient, check the buckets exist and return the ObjectStorage struct
func NewObjectStorage(endpoint string, TLS bool, accessKeyID, secretAccessKey, bucketName, bucketLocation string) (*ObjectStorage, error) {
	// Initialize minio client object
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: TLS,
	})
	if err != nil {
		return nil, terror.Error(err, "Failed to create minio client")
	}

	o := &ObjectStorage{
		Client:         minioClient,
		BucketName:     bucketName,
		BucketLocation: bucketLocation,
	}

	// Create new bucket if not exists
	err = BucketCreate(minioClient, o.BucketName, o.BucketLocation)
	if err != nil {
		return nil, terror.Error(err, "Error creating bucket")
	}

	return o, nil
}

// BucketCreate creates a new bucket if it doesn't already exist
func BucketCreate(minioClient *minio.Client, bucketName string, location string) error {
	ctx := context.Background()

	// Check if the bucket exists
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return terror.Error(err, "Failed to check if bucket exists")
	}

	// Create the bucket
	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			return terror.Error(err, "Failed to create bucket")
		}
	}

	return nil
}
