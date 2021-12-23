package nexus3

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
	"github.com/stretchr/testify/assert"
)

func TestBlobstoreGroup(t *testing.T) {
	client := getTestClient()

	bs := blobstore.File{
		Name: "test-member-name",
		Path: "test-member-path",
	}

	err := client.BlobStore.File.Create(&bs)
	assert.Nil(t, err)

	group := blobstore.Group{
		Name: "test-group",
		Members: []string{
			bs.Name,
		},
		FillPolicy: blobstore.GroupFillPolicyRoundRobin,
	}

	err = client.BlobStore.Group.Create(&group)
	assert.Nil(t, err)
	createdGroup, err := client.BlobStore.Group.Get(group.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdGroup)

	if createdGroup != nil {
		assert.Equal(t, BlobstoreGroupFillPolicyRoundRobin, createdGroup.FillPolicy)
		assert.Nil(t, createdGroup.SoftQuota)

		createdGroup.SoftQuota = &blobstore.SoftQuota{
			Type:  "spaceRemainingQuota",
			Limit: 100000000,
		}
		createdGroup.FillPolicy = BlobstoreGroupFillPolicyWriteToFirst

		err = client.BlobStore.Group.Update(createdGroup.Name, createdGroup)
		assert.Nil(t, err)

		updatedGroup, err := client.BlobStore.Group.Get(createdGroup.Name)
		assert.Nil(t, err)
		assert.NotNil(t, updatedGroup)

		if updatedGroup != nil {
			assert.NotNil(t, updatedGroup.SoftQuota)
			assert.Equal(t, BlobstoreGroupFillPolicyWriteToFirst, updatedGroup.FillPolicy)
		}

		err = client.BlobStore.Group.Delete(group.Name)
		assert.Nil(t, err)

		deletedGroup, err := client.BlobStore.Group.Get(group.Name)
		assert.NotNil(t, err)
		assert.Nil(t, deletedGroup)

		err = client.BlobStore.File.Delete(bs.Name)
		assert.Nil(t, err)

		deletedBlobstore, err := client.BlobStore.File.Get(bs.Name)
		assert.NotNil(t, err)
		assert.Nil(t, deletedBlobstore)
	}
}
