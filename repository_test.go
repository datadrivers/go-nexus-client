package client

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryRead(t *testing.T) {
	client := NewClient(getDefaultConfig())

	name := "maven-central"

	repo, err := client.RepositoryRead(name)
	assert.Nil(t, err)
	assert.NotNil(t, repo)

	if repo != nil {
		assert.Equal(t, name, repo.Name)
		assert.NotNil(t, repo.RepositoryCleanup)
		assert.NotNil(t, repo.RepositoryProxy)

		data, err := json.Marshal(repo)
		assert.Nil(t, err)
		log.Printf("%s", string(data))
	}
}

func getTestRepositoryApt(name string) Repository {
	return Repository{
		RepositoryApt: &RepositoryApt{
			Distribution: "bionic",
		},
		RepositoryAptSigning: &RepositoryAptSigning{
			Keypair:    "string",
			Passphrase: "string",
		},
		Name:   name,
		Online: true,
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}
}

func TestRepositoryAptHosted(t *testing.T) {
	client := NewClient(getDefaultConfig())
	repo := getTestRepositoryApt("test-apt-repo-hosted")

	err := client.RepositoryCreate(repo, FormatApt, TypeHosted)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryUpdate(repo.Name, updatedRepo, FormatApt, TypeHosted)
	assert.Nil(t, err)

	err = client.RepositoryDelete(repo.Name)
	assert.Nil(t, err)
}

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

	err := client.RepositoryCreate(repo, FormatDocker, TypeHosted)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryUpdate(repo.Name, updatedRepo, FormatBower, TypeHosted)
	assert.Nil(t, err)

	err = client.RepositoryDelete(repo.Name)
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

	err := client.RepositoryCreate(repo, FormatDocker, TypeHosted)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryUpdate(repo.Name, updatedRepo, FormatDocker, TypeHosted)
	assert.Nil(t, err)

	err = client.RepositoryDelete(repo.Name)
	assert.Nil(t, err)
}
