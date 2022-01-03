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

	err = service.S3.Delete(bs.Name)
	assert.Nil(t, err)
}
