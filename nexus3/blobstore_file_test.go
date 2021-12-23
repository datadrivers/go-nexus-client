package nexus3

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
	"github.com/stretchr/testify/assert"
)

func TestBlobstoreFile(t *testing.T) {
	client := getTestClient()

	bsName := "test-blobstore-name"
	bsPath := "test-blobstore-path"

	bs := blobstore.File{
		Name: bsName,
		Path: bsPath,
	}

	err := client.BlobStore.File.Create(&bs)
	assert.Nil(t, err)
	createdBlobstore, err := client.BlobStore.File.Get(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdBlobstore)

	if createdBlobstore != nil {
		assert.Equal(t, bsPath, createdBlobstore.Path)
		assert.Nil(t, createdBlobstore.SoftQuota)

		createdBlobstore.SoftQuota = &blobstore.SoftQuota{
			Type:  "spaceRemainingQuota",
			Limit: 100000000,
		}

		err = client.BlobStore.File.Update(createdBlobstore.Name, createdBlobstore)
		assert.Nil(t, err)

		updatedBlobstore, err := client.BlobStore.File.Get(createdBlobstore.Name)
		assert.Nil(t, err)
		assert.NotNil(t, updatedBlobstore)

		if updatedBlobstore != nil {
			assert.NotNil(t, updatedBlobstore.SoftQuota)
		}

		err = client.BlobStore.File.Delete(bs.Name)
		assert.Nil(t, err)

		deletedBlobstore, err := client.BlobStore.File.Get(bs.Name)
		assert.NotNil(t, err)
		assert.Nil(t, deletedBlobstore)
	}
}

func TestBlobstoreRead(t *testing.T) {
	client := getTestClient()

	bsName := "default"

	bs, err := client.BlobStore.File.Get(bsName)
	assert.Nil(t, err)
	assert.NotNil(t, bs)

	if bs != nil {
		assert.Equal(t, bsName, bs.Name)
		assert.NotEqual(t, "", bs.Path)
	}
}
