package client

import (
	"os"
	"strings"
	"testing"

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
	if os.Getenv("SKIP_S3_TESTS") != "" {
		t.Skip("Skipping S3 tests")
	}
	client := getTestClient()

	bsName := "test-blobstore-s3"
	bsType := BlobstoreTypeS3

	bs := Blobstore{
		Name: bsName,
		Type: bsType,
		BlobstoreS3BucketConfiguration: &BlobstoreS3BucketConfiguration{
			BlobstoreS3Bucket: &BlobstoreS3Bucket{
				Name:   getEnv("AWS_BUCKET_NAME", "terraform-provider-nexus-s3-test").(string),
				Region: getEnv("AWS_DEFAULT_REGION", "us-central-1").(string),
			},
			BlobstoreS3BucketSecurity: &BlobstoreS3BucketSecurity{
				AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID must be set").(string),
				SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY must be set").(string),
			},
		},
	}

	err := client.BlobstoreCreate(bs)
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

// You will need either workload identity in the target nexus or provision credentials in the nexus container to
// run this test as nexus validates the configuration when adding the blobstore
func TestBlobstoreGoogle(t *testing.T) {
	if strings.ToLower(os.Getenv("SKIP_GOOGLE_TESTS")) != "false" {
		t.Skip("Skipping Google tests")
	}
	client := getTestClient()

	bsName := "test-blobstore-google"
	bsType := BlobstoreTypeGoogle

	bs := Blobstore{
		Name: bsName,
		Type: bsType,
		BucketName: getEnv("GOOGLE_BUCKET_NAME", "terraform-provider-nexus-google-test").(string),
		Region: getEnv("GOOGLE_DEFAULT_REGION", "us-central1").(string),
		CredentialFilePath: getEnv("GOOGLE_CREDENTIAL_FILE_PATH", "").(string),
	}

	err := client.BlobstoreCreate(bs)
	assert.Nil(t, err)

	googleBS, err := client.BlobstoreRead(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, googleBS)
	if googleBS != nil {
		assert.Equal(t, BlobstoreTypeGoogle, googleBS.Type)
		assert.NotNil(t, googleBS.BucketName)
		assert.NotNil(t, googleBS.Region)
		assert.NotNil(t, googleBS.CredentialFilePath)

		err = client.BlobstoreDelete(bs.Name)
		assert.Nil(t, err)
	}
}
