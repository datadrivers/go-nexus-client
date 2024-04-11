package blobstore

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
	"github.com/stretchr/testify/assert"
)

func TestBlobstoreFile(t *testing.T) {
	service := getTestService()

	bsName := "test-blobstore-name-" + strconv.Itoa(rand.Intn(1024))
	bsPath := "test-blobstore-path"

	bs := blobstore.File{
		Name: bsName,
		Path: bsPath,
	}

	err := service.File.Create(&bs)
	assert.Nil(t, err)
	createdBlobstore, err := service.File.Get(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdBlobstore)
	assert.Equal(t, bsPath, createdBlobstore.Path)
	assert.Nil(t, createdBlobstore.SoftQuota)

	createdBlobstore.SoftQuota = &blobstore.SoftQuota{
		Type:  "spaceRemainingQuota",
		Limit: 100000000,
	}

	err = service.File.Update(createdBlobstore.Name, createdBlobstore)
	assert.Nil(t, err)

	updatedBlobstore, err := service.File.Get(createdBlobstore.Name)
	assert.Nil(t, err)
	assert.NotNil(t, updatedBlobstore)
	assert.NotNil(t, updatedBlobstore.SoftQuota)

	quotaStatus, err := getBlobstoreQuotaStatus(service.Client, updatedBlobstore.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedBlobstore.Name, quotaStatus.BlobStoreName)

	err = service.File.Delete(bs.Name)
	assert.Nil(t, err)

	deletedBlobstore, err := service.File.Get(bs.Name)
	assert.NotNil(t, err)
	assert.Nil(t, deletedBlobstore)
}

func TestBlobstoreRead(t *testing.T) {
	service := getTestService()

	bsName := "default"

	bs, err := service.File.Get(bsName)
	assert.Nil(t, err)
	assert.NotNil(t, bs)

	if bs != nil {
		assert.Equal(t, bsName, bs.Name)
		assert.NotEqual(t, "", bs.Path)
	}
}
