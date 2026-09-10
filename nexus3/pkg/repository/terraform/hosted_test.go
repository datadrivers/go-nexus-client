package terraform

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestTerraformHostedRepository(name string) repository.TerraformHostedRepository {
	writePolicy := repository.StorageWritePolicyAllow
	return repository.TerraformHostedRepository{
		Name:   name,
		Online: true,

		Storage: repository.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 &writePolicy,
		},
		TerraformSigning: repository.TerraformSigning{
			Keypair:    testSigningKeypair,
			Passphrase: tools.GetStringPointer(testSigningPassphrase),
		},
	}
}

func TestTerraformHostedRepository(t *testing.T) {
	service := getTestService()
	repo := getTestTerraformHostedRepository("test-terraform-repo-hosted-" + strconv.Itoa(rand.Intn(1024)))

	err := service.Hosted.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Hosted.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
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
