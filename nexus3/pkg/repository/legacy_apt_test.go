package repository

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestLegacyRepositoryAptHosted(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Online: true,
		Type:   repository.RepositoryTypeHosted,
		Format: repository.RepositoryFormatApt,

		Apt: &repository.AptProxy{
			Distribution: "bionic",
		},
		AptSigning: &repository.AptSigning{
			Keypair:    tools.GetStringPointer("string"),
			Passphrase: tools.GetStringPointer("string"),
		},
		Cleanup: &repository.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Storage: &repository.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 (*repository.StorageWritePolicy)(tools.GetStringPointer("ALLOW")),
		},
	}
}

func TestLegacyRepositoryAptHosted(t *testing.T) {
	service := getTestService()
	repo := getTestLegacyRepositoryAptHosted("test-apt-repo-hosted")

	err := service.Legacy.Create(repo)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Legacy.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)

	err = service.Legacy.Delete(repo.Name)
	assert.Nil(t, err)
}

func TestLegacyRepositoryAptProxy(t *testing.T) {
	service := getTestService()
	repo := getTestLegacyRepositoryAptProxy("test-repo-apt-proxy-" + strconv.Itoa(rand.Intn(1024)))

	err := service.Legacy.Create(repo)
	assert.Nil(t, err)

	createdRepo, err := service.Legacy.Get(repo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdRepo)

	if err != nil && createdRepo != nil {
		assert.Equal(t, true, createdRepo.Online)
		assert.Equal(t, repo.Name, createdRepo.Name)
		assert.Equal(t, repository.RepositoryFormatApt, createdRepo.Format)
		assert.Equal(t, repository.RepositoryTypeProxy, createdRepo.Type)
	}

	err = service.Legacy.Delete(repo.Name)
	assert.Nil(t, err)

	deletedRepo, err := service.Legacy.Get(repo.Name)
	assert.Nil(t, err)
	assert.Nil(t, deletedRepo)

}

func getTestLegacyRepositoryAptProxy(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Type:   repository.RepositoryTypeProxy,
		Format: repository.RepositoryFormatApt,
		Online: true,

		Apt: &repository.AptProxy{
			Distribution: "bionic",
			Flat:         true,
		},

		HTTPClient: &repository.HTTPClient{
			Blocked:   true,
			AutoBlock: true,
			Connection: &repository.HTTPClientConnection{
				Timeout:       tools.GetIntPointer(20),
				UseTrustStore: tools.GetBoolPointer(true),
			},
		},

		NegativeCache: &repository.NegativeCache{
			Enabled: true,
			TTL:     1440,
		},

		Proxy: &repository.Proxy{
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
			RemoteURL:      tools.GetStringPointer("https://archive.ubuntu.com/ubuntu/"),
		},

		Storage: &repository.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
	}
}
