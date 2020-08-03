package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestRepositoryDockerHostedWithPorts(name string) Repository {
	httpPort := new(int)
	httpsPort := new(int)
	*httpPort = 8082
	*httpsPort = 8083
	writePolicy := "ALLOW"

	return Repository{
		Name:   name,
		Online: true,
		Format: RepositoryFormatDocker,
		Type:   RepositoryTypeHosted,
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
			WritePolicy:                 &writePolicy,
		},
	}
}

func TestRepositoryDockerHostedWithPorts(t *testing.T) {
	client := getTestClient()
	repo := getTestRepositoryDockerHostedWithPorts("test-docker-repo-hosted-with-ports")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryUpdate(repo.Name, updatedRepo)
	assert.Nil(t, err)

	err = client.RepositoryDelete(repo.Name)
	assert.Nil(t, err)
}

func getTestRepositoryDockerHostedWithoutPorts(name string) Repository {
	writePolicy := "ALLOW_ONCE"
	return Repository{
		Name:   name,
		Online: true,
		Format: RepositoryFormatDocker,
		Type:   RepositoryTypeHosted,
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
			WritePolicy:                 &writePolicy,
		},
	}
}

func TestRepositoryDockerHostedWithoutPorts(t *testing.T) {
	client := getTestClient()
	repo := getTestRepositoryDockerHostedWithoutPorts("test-docker-repo-hosted-with-ports")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryUpdate(repo.Name, updatedRepo)
	assert.Nil(t, err)

	err = client.RepositoryDelete(repo.Name)
	assert.Nil(t, err)
}

func TestRepositoryDockerProxy(t *testing.T) {
	client := getTestClient()
	repo := getTestRepositoryDockerProxy("test-docker-repo-proxy")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	createdRepo, err := client.RepositoryRead(repo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdRepo)

	if createdRepo != nil {
		assert.Equal(t, repo.Name, createdRepo.Name)
		assert.Equal(t, repo.Type, createdRepo.Type)
		assert.Equal(t, repo.Format, createdRepo.Format)

		err := client.RepositoryDelete(repo.Name)
		assert.Nil(t, err)
	}
}

func getTestRepositoryDockerProxy(name string) Repository {
	return Repository{
		Name:   name,
		Online: true,
		Type:   RepositoryTypeProxy,
		Format: RepositoryFormatDocker,
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryDocker: &RepositoryDocker{
			V1Enabled:      false,
			ForceBasicAuth: true,
		},
		RepositoryDockerProxy: &RepositoryDockerProxy{
			IndexType: "HUB",
		},
		RepositoryHTTPClient:    &RepositoryHTTPClient{},
		RepositoryNegativeCache: &RepositoryNegativeCache{},
		RepositoryProxy: &RepositoryProxy{
			RemoteURL: "https://registry-1.docker.io",
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
		},
	}
}
