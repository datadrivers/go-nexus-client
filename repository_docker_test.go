package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestRepositoryDockerWithPorts(name string) Repository {
	httpPort := new(int)
	httpsPort := new(int)
	*httpPort = 8082
	*httpsPort = 8083

	return Repository{
		Name:   name,
		Online: true,
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryDocker: &RepositoryDocker{
			V1Enabled:      false,
			ForceBasicAuth: true,
			HTTPPort:       httpPort,
			HTTPSPort:      httpsPort,
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "ALLOW_ONCE",
		},
	}
}

func TestRepositoryDockerHostedWithPorts(t *testing.T) {
	client := NewClient(getDefaultConfig())
	repo := getTestRepositoryDockerWithPorts("test-docker-repo-hosted-with-ports")

	err := client.RepositoryDockerCreate(repo, "hosted")
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryDockerUpdate(repo.Name, updatedRepo, "hosted")
	assert.Nil(t, err)

	err = client.RepositoryDockerDelete(repo.Name)
	assert.Nil(t, err)
}

func getTestRepositoryDockerWithoutPorts(name string) Repository {
	return Repository{
		Name:   name,
		Online: true,
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryDocker: &RepositoryDocker{
			V1Enabled:      false,
			ForceBasicAuth: true,
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "ALLOW_ONCE",
		},
	}
}

func TestRepositoryDockerHostedWithoutPorts(t *testing.T) {
	client := NewClient(getDefaultConfig())
	repo := getTestRepositoryDockerWithoutPorts("test-docker-repo-hosted-with-ports")

	err := client.RepositoryDockerCreate(repo, "hosted")
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryDockerUpdate(repo.Name, updatedRepo, "hosted")
	assert.Nil(t, err)

	err = client.RepositoryDockerDelete(repo.Name)
	assert.Nil(t, err)
}
