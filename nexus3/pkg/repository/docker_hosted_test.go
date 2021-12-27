package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestDockerHostedRepository(name string) repository.DockerHostedRepository {
	writePolicy := repository.StorageWritePolicyAllow
	return repository.DockerHostedRepository{
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
		Docker: repository.Docker{
			ForceBasicAuth: true,
			V1Enabled:      false,
			HTTPPort:       tools.GetIntPointer(8180),
			HTTPSPort:      tools.GetIntPointer(8543),
		},
	}
}

func TestDockerHostedRepository(t *testing.T) {
	service := getTestService()
	repo := getTestDockerHostedRepository("test-docker-repo-hosted")

	err := service.Docker.Hosted.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Docker.Hosted.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Cleanup, generatedRepo.Cleanup)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)
	assert.Equal(t, repo.Component, generatedRepo.Component)
	assert.Equal(t, repo.Docker, generatedRepo.Docker)

	updatedRepo := repo
	updatedRepo.Online = false
	updatedRepo.Docker.V1Enabled = true

	err = service.Docker.Hosted.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Docker.Hosted.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	assert.Equal(t, updatedRepo.Docker.V1Enabled, generatedRepo.Docker.V1Enabled)

	service.Docker.Hosted.Delete(repo.Name)
	assert.Nil(t, err)
}
