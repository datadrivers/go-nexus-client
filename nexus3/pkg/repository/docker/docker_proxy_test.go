package docker

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/tools"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

func getTestDockerProxyRepository(name string) repository.DockerProxyRepository {
	return repository.DockerProxyRepository{
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
			RemoteURL:      "https://archive.ubuntu.com/ubuntu/",
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Docker: repository.Docker{
			ForceBasicAuth: true,
			V1Enabled:      false,
			HTTPPort:       tools.GetIntPointer(8280),
			HTTPSPort:      tools.GetIntPointer(8643),
		},
		DockerProxy: repository.DockerProxy{
			IndexType:                repository.DockerProxyIndexTypeHub,
			ForeignLayerUrlWhitelist: []string{},
		},
	}
}

func getTestProDockerProxyRepository(name string) repository.DockerProxyRepository {
	cacheForeignLayers := true
	return repository.DockerProxyRepository{
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
			RemoteURL:      "https://archive.ubuntu.com/ubuntu/",
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Docker: repository.Docker{
			ForceBasicAuth: true,
			V1Enabled:      false,
			HTTPPort:       tools.GetIntPointer(8280),
			HTTPSPort:      tools.GetIntPointer(8643),
			Subdomain:      tools.GetStringPointer(name),
		},
		DockerProxy: repository.DockerProxy{
			IndexType:                repository.DockerProxyIndexTypeHub,
			CacheForeignLayers:       &cacheForeignLayers,
			ForeignLayerUrlWhitelist: []string{".*\\.docker\\.io"},
		},
	}
}

func TestDockerProxyRepository(t *testing.T) {
	service := getTestService()
	repo := getTestDockerProxyRepository("test-docker-repo-hosted-" + strconv.Itoa(rand.Intn(1024)))

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
	assert.Equal(t, repo.Docker, generatedRepo.Docker)
	assert.Equal(t, repo.DockerProxy, generatedRepo.DockerProxy)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Proxy.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Proxy.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	assert.Equal(t, updatedRepo.Docker, generatedRepo.Docker)

	service.Proxy.Delete(repo.Name)
	assert.Nil(t, err)
}

func TestProDockerProxyRepository(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}
	service := getTestService()
	repo := getTestProDockerProxyRepository("test-docker-repo-hosted-" + strconv.Itoa(rand.Intn(1024)))

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
	assert.Equal(t, repo.Docker, generatedRepo.Docker)
	assert.Equal(t, repo.DockerProxy, generatedRepo.DockerProxy)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Proxy.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Proxy.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	assert.Equal(t, updatedRepo.Docker, generatedRepo.Docker)

	service.Proxy.Delete(repo.Name)
	assert.Nil(t, err)
}
