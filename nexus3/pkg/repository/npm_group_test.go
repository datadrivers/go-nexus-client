package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestNpmGroupRepository(name string) repository.NpmGroupRepository {
	return repository.NpmGroupRepository{
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

func TestNpmGroupRepository(t *testing.T) {
	service := getTestService()
	repo := getTestNpmGroupRepository("test-npm-repo-group")

	testProxyRepo := getTestNpmProxyRepository("test-npm-group-proxy")
	defer service.Npm.Proxy.Delete(testProxyRepo.Name)
	err := service.Npm.Proxy.Create(testProxyRepo)
	assert.Nil(t, err)
	repo.Group.MemberNames = append(repo.Group.MemberNames, testProxyRepo.Name)

	err = service.Npm.Group.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Npm.Group.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Group, generatedRepo.Group)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Npm.Group.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Npm.Group.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)

	service.Npm.Group.Delete(repo.Name)
	assert.Nil(t, err)
}
