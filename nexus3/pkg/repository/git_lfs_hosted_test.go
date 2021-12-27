package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestGitLfsHostedRepository(name string) repository.GitLfsHostedRepository {
	writePolicy := repository.StorageWritePolicyAllow
	return repository.GitLfsHostedRepository{
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
	}
}

func TestGitLfsHostedRepository(t *testing.T) {
	service := getTestService()
	repo := getTestGitLfsHostedRepository("test-gitlfs-repo-hosted")

	err := service.GitLfs.Hosted.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.GitLfs.Hosted.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Cleanup, generatedRepo.Cleanup)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.GitLfs.Hosted.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.GitLfs.Hosted.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)

	service.GitLfs.Hosted.Delete(repo.Name)
	assert.Nil(t, err)
}