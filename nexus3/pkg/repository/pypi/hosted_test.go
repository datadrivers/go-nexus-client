package pypi

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestPypiHostedRepository(name string) repository.PypiHostedRepository {
	writePolicy := repository.StorageWritePolicyAllow
	return repository.PypiHostedRepository{
		Name:   name,
		Online: true,

		Cleanup: &repository.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Storage: repository.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 &writePolicy,
		},
		Component: &repository.Component{
			ProprietaryComponents: true,
		},
	}
}

func TestPypiHostedRepository(t *testing.T) {
	service := getTestService()
	repo := getTestPypiHostedRepository("test-pypi-repo-hosted")

	err := service.Hosted.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Hosted.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Cleanup, generatedRepo.Cleanup)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)
	assert.Equal(t, repo.Component, generatedRepo.Component)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Hosted.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Hosted.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)

	service.Hosted.Delete(repo.Name)
	assert.Nil(t, err)
}