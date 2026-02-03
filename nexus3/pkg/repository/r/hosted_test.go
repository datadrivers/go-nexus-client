package r

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

func getTestRHostedRepository(name string) repository.RHostedRepository {
	writePolicy := repository.StorageWritePolicyAllow
	return repository.RHostedRepository{
		Name:   name,
		Online: true,

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

func TestRHostedRepository(t *testing.T) {
	service := getTestService()
	repo := getTestRHostedRepository("test-r-repo-hosted-" + strconv.Itoa(rand.Intn(1024)))

	err := service.Hosted.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Hosted.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
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
