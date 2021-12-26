package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestGoGroupRepository(name string) repository.GoGroupRepository {
	return repository.GoGroupRepository{
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

func TestGoGroupRepository(t *testing.T) {
	service := getTestService()
	repo := getTestGoGroupRepository("test-go-repo-group")

	testProxyRepo := getTestGoProxyRepository("test-go-group-proxy")
	defer service.Go.Proxy.Delete(testProxyRepo.Name)
	err := service.Go.Proxy.Create(testProxyRepo)
	assert.Nil(t, err)
	repo.Group.MemberNames = append(repo.Group.MemberNames, testProxyRepo.Name)

	err = service.Go.Group.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Go.Group.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Group, generatedRepo.Group)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Go.Group.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Go.Group.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)

	service.Go.Group.Delete(repo.Name)
	assert.Nil(t, err)
}
