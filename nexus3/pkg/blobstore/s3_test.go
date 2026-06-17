package blobstore

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
	"github.com/stretchr/testify/assert"
)

func TestBlobstoreS3(t *testing.T) {
	bucketName := tools.GetEnv("AWS_BUCKET_NAME", string("s3test-"+strconv.Itoa(rand.Intn(1024)))).(string)
	bucketLocation := tools.GetEnv("AWS_DEFAULT_REGION", "eu-central-1").(string)
	awsEndpoint := tools.GetEnv("AWS_ENDPOINT", string("")).(string)
	awsAccessKeyID := tools.GetEnv("AWS_ACCESS_KEY_ID", string("")).(string)
	awsSecretAccessKey := tools.GetEnv("AWS_SECRET_ACCESS_KEY", string("")).(string)

	service := getTestService()

	bsName := "test-blobstore-s3"
	forcePathStyle := true

	bs := &blobstore.S3{
		Name: bsName,
		BucketConfiguration: blobstore.S3BucketConfiguration{
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

	err := service.S3.Create(bs)
	assert.Nil(t, err)

	s3BS, err := service.S3.Get(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, s3BS)
	assert.NotNil(t, s3BS.BucketConfiguration)
	assert.NotNil(t, s3BS.BucketConfiguration.Bucket)
	assert.NotNil(t, s3BS.BucketConfiguration.BucketSecurity)

	s3BS.SoftQuota = &blobstore.SoftQuota{
		Type:  "spaceRemainingQuota",
		Limit: 100000000,
	}

	err = service.S3.Update(s3BS.Name, s3BS)
	assert.Nil(t, err)

	updatedBlobstore, err := service.S3.Get(s3BS.Name)
	assert.Nil(t, err)
	assert.NotNil(t, updatedBlobstore)
	assert.NotNil(t, updatedBlobstore.SoftQuota)

	quotaStatus, err := getBlobstoreQuotaStatus(service.Client, updatedBlobstore.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedBlobstore.Name, quotaStatus.BlobStoreName)

	err = service.S3.Delete(bs.Name)
	assert.Nil(t, err)
}

func TestBlobStoreS3PresignedUrl(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}
	bucketName := tools.GetEnv("AWS_BUCKET_NAME", string("s3test-"+strconv.Itoa(rand.Intn(1024)))).(string)
	bucketLocation := tools.GetEnv("AWS_DEFAULT_REGION", "eu-central-1").(string)
	awsEndpoint := tools.GetEnv("AWS_ENDPOINT", string("")).(string)
	awsAccessKeyID := tools.GetEnv("AWS_ACCESS_KEY_ID", string("")).(string)
	awsSecretAccessKey := tools.GetEnv("AWS_SECRET_ACCESS_KEY", string("")).(string)

	service := getTestService()

	bsName := "test-blobstore-s3-presigned-url"
	forcePathStyle := true
	preSignedUrlEnabled := true

	bs := &blobstore.S3{
		Name: bsName,
		BucketConfiguration: blobstore.S3BucketConfiguration{
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
			PreSignedUrlEnabled: &preSignedUrlEnabled,
		},
	}

	err := service.S3.Create(bs)
	assert.Nil(t, err)

	s3BS, err := service.S3.Get(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, s3BS)
	assert.NotNil(t, s3BS.BucketConfiguration)
	assert.NotNil(t, s3BS.BucketConfiguration.Bucket)
	assert.NotNil(t, s3BS.BucketConfiguration.BucketSecurity)
	assert.True(t, *s3BS.BucketConfiguration.PreSignedUrlEnabled)

	err = service.S3.Delete(bs.Name)
	assert.Nil(t, err)
}
