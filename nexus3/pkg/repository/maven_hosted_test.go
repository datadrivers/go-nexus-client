package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestMavenHostedRepository(name string) repository.MavenHostedRepository {
	writePolicy := repository.StorageWritePolicyAllow
	versionPolicy := repository.MavenVersionPolicySnapshot
	layoutPolicy := repository.MavenLayoutPolicyStrict
	return repository.MavenHostedRepository{
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
		Maven: repository.Maven{
			VersionPolicy: &versionPolicy,
			LayoutPolicy:  &layoutPolicy,
		},
	}
}

func TestMavenHostedRepository(t *testing.T) {
	service := getTestService()
	repo := getTestMavenHostedRepository("test-maven-repo-hosted")

	err := service.Maven.Hosted.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Maven.Hosted.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Cleanup, generatedRepo.Cleanup)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)
	assert.Equal(t, repo.Maven, generatedRepo.Maven)

	updatedRepo := repo
	updatedRepo.Online = false
	newVersionPolicy := repository.MavenVersionPolicyMixed
	updatedRepo.Maven.VersionPolicy = &newVersionPolicy

	err = service.Maven.Hosted.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Maven.Hosted.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	assert.Equal(t, updatedRepo.Maven.VersionPolicy, generatedRepo.Maven.VersionPolicy)

	service.Maven.Hosted.Delete(repo.Name)
	assert.Nil(t, err)
}
