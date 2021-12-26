package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestBowerGroupRepository(name string) repository.BowerGroupRepository {
	return repository.BowerGroupRepository{
		Name:   name,
		Online: true,

		Group: repository.Group{
			MemberNames: []string{},
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
	}
}

func TestBowerGroupRepository(t *testing.T) {
	service := getTestService()
	repo := getTestBowerGroupRepository("test-bower-repo-group")

	testProxyRepo := getTestBowerProxyRepository("test-bower-group-proxy")
	defer service.Bower.Proxy.Delete(testProxyRepo.Name)
	err := service.Bower.Proxy.Create(testProxyRepo)
	assert.Nil(t, err)
	repo.Group.MemberNames = append(repo.Group.MemberNames, testProxyRepo.Name)

	err = service.Bower.Group.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Bower.Group.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Group, generatedRepo.Group)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Bower.Group.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Bower.Group.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)

	service.Bower.Group.Delete(repo.Name)
	assert.Nil(t, err)
}
