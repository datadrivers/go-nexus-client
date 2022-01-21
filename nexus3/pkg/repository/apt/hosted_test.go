package apt_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestAptHostedRepository(name string) repository.AptHostedRepository {
	writePolicy := repository.StorageWritePolicyAllow
	return repository.AptHostedRepository{
		Name:   name,
		Online: true,

		Apt: repository.AptHosted{
			Distribution: "bionic",
		},
		AptSigning: repository.AptSigning{
			Keypair:    "string",
			Passphrase: tools.GetStringPointer("string"),
		},
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

func TestAptHostedRepository(t *testing.T) {
	service := getTestService()
	repo := getTestAptHostedRepository("test-apt-repo-hosted-" + strconv.Itoa(rand.Intn(1024)))

	err := service.Hosted.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Hosted.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Apt, generatedRepo.Apt)
	assert.Equal(t, repo.Cleanup, generatedRepo.Cleanup)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)

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
