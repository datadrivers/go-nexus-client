package yum

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestYumGroupRepository(name string) repository.YumGroupRepository {
	return repository.YumGroupRepository{
		Name:   name,
		Online: true,

		Group: repository.Group{
			MemberNames: []string{},
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		YumSigning: &repository.YumSigning{
			Keypair:    tools.GetStringPointer("test-key"),
			Passphrase: tools.GetStringPointer("test"),
		},
	}
}

func TestYumGroupRepository(t *testing.T) {
	service := getTestService()
	repo := getTestYumGroupRepository("test-yum-repo-group-" + strconv.Itoa(rand.Intn(1024)))

	testProxyRepo := getTestYumProxyRepository("test-yum-group-proxy-" + strconv.Itoa(rand.Intn(1024)))
	defer service.Proxy.Delete(testProxyRepo.Name)
	err := service.Proxy.Create(testProxyRepo)
	assert.Nil(t, err)
	repo.Group.MemberNames = append(repo.Group.MemberNames, testProxyRepo.Name)

	err = service.Group.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Group.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Group, generatedRepo.Group)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)
	// ToDo: Add the following test after the implementation of this issue https://issues.sonatype.org/browse/NEXUS-30751
	// assert.Equal(t, repo.YumSigning, generatedRepo.YumSigning)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Group.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Group.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)

	service.Group.Delete(repo.Name)
	assert.Nil(t, err)
}
