package client

import (
	"testing"

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

	bsCreated, err := client.BlobstoreRead(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, bsCreated)
	// Path not returned by API, not possible to test :-/
	// assert.Equal(t, bsPath, bsCreated.Path)
	assert.Equal(t, bsType, bsCreated.Type)
	assert.Equal(t, 0, bsCreated.BlobCount)
	assert.Nil(t, bsCreated.BlobstoreSoftQuota)

	bsCreated.BlobstoreSoftQuota = &BlobstoreSoftQuota{
		Type:  "spaceRemainingQuota",
		Limit: 100000000,
	}
	err = client.BlobstoreUpdate(bsCreated.Name, *bsCreated)
	assert.Nil(t, err)

	bsUpdated, err := client.BlobstoreRead(bsCreated.Name)
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

	bs, err := client.BlobstoreRead(bsName)
	assert.Nil(t, err)
	assert.NotNil(t, bs)

	if bs != nil {
		assert.Equal(t, bsName, bs.Name)
	}
}
