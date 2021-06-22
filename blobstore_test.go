package client

import (
	"context"
	"testing"

	minio "github.com/minio/minio-go/v7"
	credentials "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/assert"
)

func TestBlobstoreFile(t *testing.T) {
	client := getTestClient()

	bsName := "test-blobstore-name"
	bsPath := "test-blobstore-path"
	bsType := BlobstoreTypeFile

	bs := Blobstore{
		Name: bsName,
		Path: bsPath,
		Type: bsType,
	}

	err := client.BlobstoreCreate(bs)
	assert.Nil(t, err)

	createdBlobstore, err := client.BlobstoreRead(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdBlobstore)

	if createdBlobstore != nil {
		assert.Equal(t, bsPath, createdBlobstore.Path)
		assert.Equal(t, bsType, createdBlobstore.Type)
		assert.Equal(t, 0, createdBlobstore.BlobCount)
		assert.Nil(t, createdBlobstore.BlobstoreSoftQuota)

		createdBlobstore.BlobstoreSoftQuota = &BlobstoreSoftQuota{
			Type:  "spaceRemainingQuota",
			Limit: 100000000,
		}

		err = client.BlobstoreUpdate(createdBlobstore.Name, *createdBlobstore)
		assert.Nil(t, err)

		updatedBlobstore, err := client.BlobstoreRead(createdBlobstore.Name)
		assert.Nil(t, err)
		assert.NotNil(t, updatedBlobstore)

		if updatedBlobstore != nil {
			assert.NotNil(t, updatedBlobstore.BlobstoreSoftQuota)
		}

		err = client.BlobstoreDelete(bs.Name)
		assert.Nil(t, err)

		deletedBlobstore, err := client.BlobstoreRead(bs.Name)
		assert.Nil(t, err)
		assert.Nil(t, deletedBlobstore)
	}
}

func TestBlobstoreRead(t *testing.T) {
	client := getTestClient()

	bsName := "default"

	bs, err := client.BlobstoreRead(bsName)
	assert.Nil(t, err)
	assert.NotNil(t, bs)

	if bs != nil {
		assert.Equal(t, bsName, bs.Name)
		assert.NotEqual(t, "", bs.Path)
	}
}

func TestBlobstoreS3(t *testing.T) {
	bucketName := "s3test"
	bucketLocation := "us-east-1"
	minioGoEndpoint := "localhost:9000"
	minioNexusEndpoint := "http://minio:9000"
	minioAccessKeyID := "minioadmin"
	minioSecretAccessKey := "minioadmin"
	minioUseSSL := false

	err := ensureMinioBucket(bucketName, bucketLocation, minioGoEndpoint, minioUseSSL, minioAccessKeyID, minioSecretAccessKey)
	assert.Nil(t, err)

	client := getTestClient()

	bsName := "test-blobstore-s3"
	bsType := BlobstoreTypeS3

	bs := Blobstore{
		Name: bsName,
		Type: bsType,
		BlobstoreS3BucketConfiguration: &BlobstoreS3BucketConfiguration{
			BlobstoreS3Bucket: &BlobstoreS3Bucket{
				Name:   bucketName,
				Region: bucketLocation,
			},
			BlobstoreS3BucketSecurity: &BlobstoreS3BucketSecurity{
				AccessKeyID:     minioAccessKeyID,
				SecretAccessKey: minioSecretAccessKey,
			},
			BlobstoreS3AdvancedBucketConnection: &BlobstoreS3AdvancedBucketConnection{
				Endpoint:       minioNexusEndpoint,
				ForcePathStyle: true,
			},
		},
	}

	err = client.BlobstoreCreate(bs)
	assert.Nil(t, err)

	s3BS, err := client.BlobstoreRead(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, s3BS)
	if s3BS != nil {
		assert.Equal(t, BlobstoreTypeS3, s3BS.Type)
		assert.NotNil(t, s3BS.BlobstoreS3BucketConfiguration)
		assert.NotNil(t, s3BS.BlobstoreS3BucketConfiguration.BlobstoreS3Bucket)
		assert.NotNil(t, s3BS.BlobstoreS3BucketConfiguration.BlobstoreS3BucketSecurity)

		err = client.BlobstoreDelete(bs.Name)
		assert.Nil(t, err)
	}
}

func ensureMinioBucket(bucketName string, bucketLocation string, endpoint string, useSSL bool, accessKeyID string, secretAccessKey string) error {

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		return err
	}
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: bucketLocation})
	}
	return err
}
