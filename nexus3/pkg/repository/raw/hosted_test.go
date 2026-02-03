package raw

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

func getTestRawHostedRepository(name string) repository.RawHostedRepository {
	writePolicy := repository.StorageWritePolicyAllow
	contentDisposition := repository.RawContentDispositionAttachment
	return repository.RawHostedRepository{
		Name:   name,
		Online: true,

		Storage: repository.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 &writePolicy,
		},
		Raw: &repository.Raw{
			ContentDisposition: &contentDisposition,
		},
	}
}

func TestRawHostedRepository(t *testing.T) {
	service := getTestService()
	repo := getTestRawHostedRepository("test-raw-repo-hosted-" + strconv.Itoa(rand.Intn(1024)))

	err := service.Hosted.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Hosted.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)
	// ToDo: Add following Test after implemented issue https://issues.sonatype.org/browse/NEXUS-30750
	// assert.Equal(t, repo.Raw, generatedRepo.Raw)

	newContentDisposition := repository.RawContentDispositionInline
	updatedRepo := repo
	updatedRepo.Online = false
	updatedRepo.Raw.ContentDisposition = &newContentDisposition

	err = service.Hosted.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Hosted.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	// ToDo: Add following Test after implemented issue https://issues.sonatype.org/browse/NEXUS-30750
	// assert.Equal(t, updatedRepo.Raw, generatedRepo.Raw)

	service.Hosted.Delete(repo.Name)
	assert.Nil(t, err)
}
