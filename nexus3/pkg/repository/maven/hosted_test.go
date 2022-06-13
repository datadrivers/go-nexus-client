package maven

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestMavenHostedRepository(name string) repository.MavenHostedRepository {
	writePolicy := repository.StorageWritePolicyAllow
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
			VersionPolicy: repository.MavenVersionPolicySnapshot,
			LayoutPolicy:  repository.MavenLayoutPolicyStrict,
		},
	}
}

func TestMavenHostedRepository(t *testing.T) {
	service := getTestService()
	repo := getTestMavenHostedRepository("test-maven-repo-hosted-" + strconv.Itoa(rand.Intn(1024)))

	err := service.Hosted.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Hosted.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Cleanup, generatedRepo.Cleanup)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)
	assert.Equal(t, repo.Maven, generatedRepo.Maven)

	updatedRepo := repo
	updatedRepo.Online = false
	updatedRepo.VersionPolicy = repository.MavenVersionPolicyMixed

	err = service.Hosted.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Hosted.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	assert.Equal(t, updatedRepo.VersionPolicy, generatedRepo.VersionPolicy)

	service.Hosted.Delete(repo.Name)
	assert.Nil(t, err)
}
