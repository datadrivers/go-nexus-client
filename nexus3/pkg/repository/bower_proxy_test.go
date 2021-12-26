package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestBowerProxyRepository(name string) repository.BowerProxyRepository {
	return repository.BowerProxyRepository{
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
			RemoteURL:      tools.GetStringPointer("https://archive.ubuntu.com/ubuntu/"),
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},

		Bower: &repository.Bower{
			RewritePackageUrls: true,
		},
	}
}

func TestBowerProxyRepository(t *testing.T) {
	service := getTestService()
	repo := getTestBowerProxyRepository("test-bower-repo-hosted")

	err := service.Bower.Proxy.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Bower.Proxy.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.HTTPClient.Blocked, generatedRepo.HTTPClient.Blocked)
	assert.Equal(t, repo.HTTPClient.AutoBlock, generatedRepo.HTTPClient.AutoBlock)
	assert.Equal(t, repo.HTTPClient.Connection.Timeout, generatedRepo.HTTPClient.Connection.Timeout)
	assert.Equal(t, repo.HTTPClient.Connection.UseTrustStore, generatedRepo.HTTPClient.Connection.UseTrustStore)
	assert.Equal(t, repo.NegativeCache, generatedRepo.NegativeCache)
	assert.Equal(t, repo.Proxy, generatedRepo.Proxy)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)
	assert.Equal(t, repo.Bower, generatedRepo.Bower)

	updatedRepo := repo
	updatedRepo.Online = false
	updatedRepo.Bower.RewritePackageUrls = false

	err = service.Bower.Proxy.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Bower.Proxy.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	assert.Equal(t, updatedRepo.Bower, generatedRepo.Bower)

	service.Bower.Proxy.Delete(repo.Name)
	assert.Nil(t, err)
}
