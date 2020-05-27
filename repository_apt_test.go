package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestRepositoryAptHosted(name string) Repository {
	return Repository{
		Name:   name,
		Online: true,
		Type:   RepositoryTypeHosted,
		Format: RepositoryFormatApt,

		RepositoryApt: &RepositoryApt{
			Distribution: "bionic",
		},
		RepositoryAptSigning: &RepositoryAptSigning{
			Keypair:    "string",
			Passphrase: "string",
		},
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "ALLOW",
		},
	}
}

func TestRepositoryAptHosted(t *testing.T) {
	client := NewClient(getDefaultConfig())
	repo := getTestRepositoryAptHosted("test-apt-repo-hosted")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryUpdate(repo.Name, updatedRepo)
	assert.Nil(t, err)

	err = client.RepositoryDelete(repo.Name)
	assert.Nil(t, err)
}

func TestRepositoryAptProxy(t *testing.T) {
	client := NewClient(getDefaultConfig())
	repo := getTestRepositoryAptProxy("test-repo-apt-proxy")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	if err != nil {
		createdRepo, err := client.RepositoryRead(repo.Name)
		assert.Nil(t, err)
		assert.NotNil(t, createdRepo)

		if err != nil && createdRepo != nil {
			assert.Equal(t, true, createdRepo.Online)
			assert.Equal(t, repo.Name, createdRepo.Name)
			assert.Equal(t, RepositoryFormatApt, createdRepo.Format)
			assert.Equal(t, RepositoryTypeProxy, createdRepo.Type)
		}

		err = client.RepositoryDelete(repo.Name)
		assert.Nil(t, err)
	}
}

func getTestRepositoryAptProxy(name string) Repository {
	return Repository{
		Name:   name,
		Type:   RepositoryTypeProxy,
		Format: RepositoryFormatApt,
		Online: true,

		RepositoryApt: &RepositoryApt{
			Distribution: "bionic",
			Flat:         true,
		},

		RepositoryHTTPClient: &RepositoryHTTPClient{
			Blocked:   true,
			AutoBlock: true,
		},

		RepositoryNegativeCache: &RepositoryNegativeCache{
			Enabled: true,
			TTL:     1440,
		},

		RepositoryProxy: &RepositoryProxy{
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
			RemoteURL:      "http://archive.ubuntu.com/ubuntu/",
		},

		RepositoryStorage: &RepositoryStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
	}
}
