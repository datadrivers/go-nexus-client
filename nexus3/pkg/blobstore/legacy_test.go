package blobstore

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
	"github.com/stretchr/testify/assert"
)

func TestLegacyBlobstoreFile(t *testing.T) {
	service := getTestService()

	bsName := "test-blobstore-name"
	bsPath := "test-blobstore-path"
	bsType := blobstore.BlobstoreTypeFile

	bs := blobstore.Legacy{
		Name: bsName,
		Path: bsPath,
		Type: bsType,
	}

	err := service.Legacy.Create(&bs)
	assert.Nil(t, err)

	createdBlobstore, err := service.Legacy.Get(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdBlobstore)

	if createdBlobstore != nil {
		assert.Equal(t, bsPath, createdBlobstore.Path)
		assert.Equal(t, bsType, createdBlobstore.Type)
		assert.Equal(t, 0, createdBlobstore.BlobCount)
		assert.Nil(t, createdBlobstore.SoftQuota)

		createdBlobstore.SoftQuota = &blobstore.SoftQuota{
			Type:  "spaceRemainingQuota",
			Limit: 100000000,
		}

		err = service.Legacy.Update(createdBlobstore.Name, *createdBlobstore)
		assert.Nil(t, err)

		updatedBlobstore, err := service.Legacy.Get(createdBlobstore.Name)
		assert.Nil(t, err)
		assert.NotNil(t, updatedBlobstore)

		if updatedBlobstore != nil {
			assert.NotNil(t, updatedBlobstore.SoftQuota)
		}

		err = service.Legacy.Delete(bs.Name)
		assert.Nil(t, err)

		deletedBlobstore, err := service.Legacy.Get(bs.Name)
		assert.Nil(t, err)
		assert.Nil(t, deletedBlobstore)
	}
}

func TestLegacyBlobstoreRead(t *testing.T) {
	service := getTestService()

	bsName := "default"

	bs, err := service.Legacy.Get(bsName)
	assert.Nil(t, err)
	assert.NotNil(t, bs)

	if bs != nil {
		assert.Equal(t, bsName, bs.Name)
		assert.NotEqual(t, "", bs.Path)
	}
}

func TestLegacyBlobstoreS3(t *testing.T) {
	bucketName := tools.GetEnv("AWS_BUCKET_NAME", string("s3test-"+strconv.Itoa(rand.Intn(1024)))).(string)
	bucketLocation := tools.GetEnv("AWS_DEFAULT_REGION", "eu-central-1").(string)
	awsEndpoint := tools.GetEnv("AWS_ENDPOINT", string("")).(string)
	awsAccessKeyID := tools.GetEnv("AWS_ACCESS_KEY_ID", string("")).(string)
	awsSecretAccessKey := tools.GetEnv("AWS_SECRET_ACCESS_KEY", string("")).(string)
	service := getTestService()

	bsName := "test-blobstore-s3"
	bsType := blobstore.BlobstoreTypeS3
	forcePathStyle := true

	bs := blobstore.Legacy{
		Name: bsName,
		Type: bsType,
		S3BucketConfiguration: &blobstore.S3BucketConfiguration{
			Bucket: blobstore.S3Bucket{
				Name:   bucketName,
				Region: bucketLocation,
			},
			BucketSecurity: &blobstore.S3BucketSecurity{
				AccessKeyID:     awsAccessKeyID,
				SecretAccessKey: awsSecretAccessKey,
			},
			AdvancedBucketConnection: &blobstore.S3AdvancedBucketConnection{
				Endpoint:       awsEndpoint,
				ForcePathStyle: &forcePathStyle,
			},
		},
	}

	err := service.Legacy.Create(&bs)
	assert.Nil(t, err)

	s3BS, err := service.Legacy.Get(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, s3BS)
	if s3BS != nil {
		assert.Equal(t, blobstore.BlobstoreTypeS3, s3BS.Type)
		assert.NotNil(t, s3BS.S3BucketConfiguration)
		assert.NotNil(t, s3BS.S3BucketConfiguration.Bucket)
		assert.NotNil(t, s3BS.S3BucketConfiguration.BucketSecurity)

		err = service.Legacy.Delete(bs.Name)
		assert.Nil(t, err)
	}
}
