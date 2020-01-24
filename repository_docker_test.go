package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestRepositoryDocker(name string) Repository {
	return Repository{
		Name:   name,
		Online: true,
		Cleanup: RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Docker: &Docker{
			V1Enabled:      false,
			ForceBasicAuth: true,
			HTTPPort:       8082,
			HTTPSPort:      8083,
		},
		Storage: RepositoryStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}
}

func TestRepositoryDockerHosted(t *testing.T) {
	client := NewClient(getDefaultConfig())
	repo := getTestRepositoryDocker("test-docker-repo-hosted")

	err := client.RepositoryDockerCreate(repo, "hosted")
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryDockerUpdate(repo.Name, updatedRepo, "hosted")
	assert.Nil(t, err)

	err = client.RepositoryDockerDelete(repo.Name)
	assert.Nil(t, err)
}
