package client

import (
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func TestBlobstoreFile(t *testing.T) {
	client := NewClient(getDefaultConfig())

	bsName := "test-blobstore-name"
	bsPath := "test-blobstore-path"
	bsType := BlobstoreTypeFile

	bs := Blobstore{
		Name: bsName,
		Path: bsPath,
		Type: bsType,
	}

	createErr := client.BlobstoreCreate(bs)
	assert.Nil(t, createErr)

	bsCreated, err := client.BlobstoreReadSpecified(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, bsCreated)
	assert.Equal(t, bsPath, bsCreated.Path)
	assert.Equal(t, bsType, bsCreated.Type)
	assert.Equal(t, 0, bsCreated.BlobCount)
	assert.Nil(t, bsCreated.BlobstoreSoftQuota)

	bsCreated.BlobstoreSoftQuota = &BlobstoreSoftQuota{
		Type:  "spaceRemainingQuota",
		Limit: 100000000,
	}
	err = client.BlobstoreUpdate(bsCreated.Name, *bsCreated)
	assert.Nil(t, err)

	bsUpdated, err := client.BlobstoreReadSpecified(bsCreated.Name)
	assert.Nil(t, err)
	assert.NotNil(t, bsUpdated)
	assert.NotNil(t, bsUpdated.BlobstoreSoftQuota)

	if createErr == nil {
		err := client.BlobstoreDelete(bs.Name)
		assert.Nil(t, err)
	}
}

func TestBlobstoreRead(t *testing.T) {
	client := NewClient(getDefaultConfig())

	bsName := "default"

	bs, err := client.BlobstoreReadSpecified(bsName)
	assert.Nil(t, err)
	assert.NotNil(t, bs)

	if bs != nil {
		assert.Equal(t, bsName, bs.Name)
		assert.NotEqual(t, "", bs.Path)
	}
}

func TestBlobstoreS3(t *testing.T) {
	client := NewClient(getDefaultConfig())

	bsName := "test-blobstore-s3"
	bsType := BlobstoreTypeS3
	var forcePathStyle, _ = strconv.ParseBool(getEnv("AWS_FORCE_PATH", "true"))
	var expiration, _ = strconv.Atoi(getEnv("AWS_EXPIRATION", "1"))

	bs := Blobstore{
		Name: bsName,
		Type: bsType,
		BlobstoreS3BucketConfiguration: &BlobstoreS3BucketConfiguration{
			BlobstoreS3Bucket: &BlobstoreS3Bucket{
				Name:   	getEnv("AWS_BUCKET_NAME", "test"),
				Region: 	getEnv("AWS_DEFAULT_REGION", "eu-central-1"),
				Expiration: expiration,
			},
			BlobstoreS3BucketSecurity: &BlobstoreS3BucketSecurity{
				AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", "minioadmin"),
				SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", "minioadmin"),
			},
			BlobstoreS3AdvancedBucketConnection: &BlobstoreS3AdvancedBucketConnection{
				Endpoint:       getEnv("AWS_ENDPOINT", "http://minio:9000"),
				ForcePathStyle: forcePathStyle,
			},
		},
	}

	err := client.BlobstoreCreate(bs)
	assert.Nil(t, err)

	s3BS, err := client.BlobstoreReadSpecified(bs.Name)
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
