package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		Cleanup: RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Storage: RepositoryStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "allow_once",
		},
	}
}

func TestRepositoryAptHosted(t *testing.T) {
	client := NewClient(getDefaultConfig())
	repo := getTestRepositoryApt("test-apt-repo-hosted")

	err := client.RepositoryAptCreate(repo, "hosted")
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryAptUpdate(repo.Name, updatedRepo, "hosted")
	assert.Nil(t, err)

	err = client.RepositoryAptDelete(repo.Name)
	assert.Nil(t, err)
}
