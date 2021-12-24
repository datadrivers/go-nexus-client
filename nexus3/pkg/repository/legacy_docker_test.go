package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestLegacyRepositoryDockerHostedWithPorts(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Online: true,
		Format: repository.RepositoryFormatDocker,
		Type:   repository.RepositoryTypeHosted,
		Cleanup: &repository.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Docker: &repository.Docker{
			V1Enabled:      false,
			ForceBasicAuth: true,
			HTTPPort:       tools.GetIntPointer(8082),
			HTTPSPort:      tools.GetIntPointer(8083),
		},
		Storage: &repository.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 tools.GetStringPointer("ALLOW"),
		},
	}
}

func TestLegacyRepositoryDockerHostedWithPorts(t *testing.T) {
	service := getTestService()
	repo := getTestLegacyRepositoryDockerHostedWithPorts("test-docker-repo-hosted-with-ports")

	err := service.Legacy.Create(repo)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Legacy.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)

	err = service.Legacy.Delete(repo.Name)
	assert.Nil(t, err)
}

func getTestLegacyRepositoryDockerHostedWithoutPorts(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Online: true,
		Format: repository.RepositoryFormatDocker,
		Type:   repository.RepositoryTypeHosted,
		Cleanup: &repository.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Docker: &repository.Docker{
			V1Enabled:      false,
			ForceBasicAuth: true,
		},
		Storage: &repository.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 tools.GetStringPointer("ALLOW_ONCE"),
		},
	}
}

func TestLegacyRepositoryDockerHostedWithoutPorts(t *testing.T) {
	service := getTestService()
	repo := getTestLegacyRepositoryDockerHostedWithoutPorts("test-docker-repo-hosted-with-ports")

	err := service.Legacy.Create(repo)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Legacy.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)

	err = service.Legacy.Delete(repo.Name)
	assert.Nil(t, err)
}

func TestLegacyRepositoryDockerProxy(t *testing.T) {
	service := getTestService()
	repo := getTestLegacyRepositoryDockerProxy("test-docker-repo-proxy")

	err := service.Legacy.Create(repo)
	assert.Nil(t, err)

	createdRepo, err := service.Legacy.Get(repo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdRepo)

	if createdRepo != nil {
		assert.Equal(t, repo.Name, createdRepo.Name)
		assert.Equal(t, repo.Type, createdRepo.Type)
		assert.Equal(t, repo.Format, createdRepo.Format)

		err := service.Legacy.Delete(repo.Name)
		assert.Nil(t, err)
	}
}

func getTestLegacyRepositoryDockerProxy(name string) repository.LegacyRepository {
	dockerIndexType := repository.DockerProxyIndexTypeHub
	return repository.LegacyRepository{
		Name:   name,
		Online: true,
		Type:   repository.RepositoryTypeProxy,
		Format: repository.RepositoryFormatDocker,
		Cleanup: &repository.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Docker: &repository.Docker{
			V1Enabled:      false,
			ForceBasicAuth: true,
		},
		DockerProxy: &repository.DockerProxy{
			IndexType: &dockerIndexType,
		},
		HTTPClient: &repository.HTTPClient{
			Connection: &repository.HTTPClientConnection{
				Timeout: tools.GetIntPointer(20),
			},
		},
		NegativeCache: &repository.NegativeCache{},
		Proxy: &repository.Proxy{
			RemoteURL: tools.GetStringPointer("https://registry-1.docker.io"),
		},
		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
		},
	}
}
