package blobstore

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
	"github.com/stretchr/testify/assert"
)

func TestBlobstoreGroup(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus blobstore for Azure tests")
	}

	service := getTestService()

	bs := blobstore.File{
		Name: "test-member-name-" + strconv.Itoa(rand.Intn(1024)),
		Path: "test-member-path",
	}

	err := service.File.Create(&bs)
	assert.Nil(t, err)

	group := blobstore.Group{
		Name: "test-group",
		Members: []string{
			bs.Name,
		},
		FillPolicy: blobstore.GroupFillPolicyRoundRobin,
	}

	err = service.Group.Create(&group)
	assert.Nil(t, err)
	createdGroup, err := service.Group.Get(group.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdGroup)

	assert.Equal(t, blobstore.GroupFillPolicyRoundRobin, createdGroup.FillPolicy)
	assert.Nil(t, createdGroup.SoftQuota)

	createdGroup.SoftQuota = &blobstore.SoftQuota{
		Type:  "spaceRemainingQuota",
		Limit: 100000000,
	}
	createdGroup.FillPolicy = blobstore.GroupFillPolicyWriteToFirst

	err = service.Group.Update(createdGroup.Name, createdGroup)
	assert.Nil(t, err)

	updatedGroup, err := service.Group.Get(createdGroup.Name)
	assert.Nil(t, err)
	assert.NotNil(t, updatedGroup)

	assert.NotNil(t, updatedGroup.SoftQuota)
	assert.Equal(t, blobstore.GroupFillPolicyWriteToFirst, updatedGroup.FillPolicy)

	err = service.Group.Delete(group.Name)
	assert.Nil(t, err)

	deletedGroup, err := service.Group.Get(group.Name)
	assert.NotNil(t, err)
	assert.Nil(t, deletedGroup)

	err = service.File.Delete(bs.Name)
	assert.Nil(t, err)

	deletedBlobstore, err := service.File.Get(bs.Name)
	assert.NotNil(t, err)
	assert.Nil(t, deletedBlobstore)

}
