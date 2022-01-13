package pypi

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestPypiProxyRepository(name string) repository.PypiProxyRepository {
	return repository.PypiProxyRepository{
		Name:   name,
		Online: true,
		HTTPClient: repository.HTTPClient{
			Blocked:   true,
			AutoBlock: true,
			Connection: &repository.HTTPClientConnection{
				Timeout:       tools.GetIntPointer(20),
				UseTrustStore: tools.GetBoolPointer(true),
			},
		},
		NegativeCache: repository.NegativeCache{
			Enabled: true,
			TTL:     1440,
		},
		Proxy: repository.Proxy{
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
			RemoteURL:      "https://api.pypi.org/v3/index.json",
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
	}
}

func TestPypiProxyRepository(t *testing.T) {
	service := getTestService()
	repo := getTestPypiProxyRepository("test-pypi-repo-hosted-" + strconv.Itoa(rand.Intn(1024)))

	err := service.Proxy.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Proxy.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.HTTPClient.Blocked, generatedRepo.HTTPClient.Blocked)
	assert.Equal(t, repo.HTTPClient.AutoBlock, generatedRepo.HTTPClient.AutoBlock)
	assert.Equal(t, repo.HTTPClient.Connection.Timeout, generatedRepo.HTTPClient.Connection.Timeout)
	assert.Equal(t, repo.HTTPClient.Connection.UseTrustStore, generatedRepo.HTTPClient.Connection.UseTrustStore)
	assert.Equal(t, repo.NegativeCache, generatedRepo.NegativeCache)
	assert.Equal(t, repo.Proxy, generatedRepo.Proxy)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)

	updatedRepo := repo
	updatedRepo.Online = false
	updatedRepo.Proxy.RemoteURL = "https://api.pypi.org/v2/"

	err = service.Proxy.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Proxy.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)

	service.Proxy.Delete(repo.Name)
	assert.Nil(t, err)
}
