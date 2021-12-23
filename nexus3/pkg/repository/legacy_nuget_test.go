package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func TestLegacyRepositoryNugetProxy(t *testing.T) {
	service := getTestService()
	repo := getTestLegacyRepositoryNugetProxy("test-nuget-proxy-repo")

	err := service.Legacy.Create(repo)
	assert.Nil(t, err)

	createdRepo, err := service.Legacy.Get(repo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdRepo)

	if createdRepo != nil {

		err := service.Legacy.Delete(createdRepo.Name)
		assert.Nil(t, err)
	}
}

func getTestLegacyRepositoryNugetProxy(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Format: repository.RepositoryFormatNuget,
		Name:   name,
		Online: true,
		Type:   repository.RepositoryTypeProxy,

		Cleanup: &repository.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		HTTPClient: &repository.HTTPClient{
			Connection: &repository.HTTPClientConnection{
				Timeout: tools.GetIntPointer(20),
			},
		},
		NegativeCache: &repository.NegativeCache{},
		NugetProxy: &repository.NugetProxy{
			QueryCacheItemMaxAge: 1,
		},
		Proxy: &repository.Proxy{
			RemoteURL: "https://www.nuget.org/api/v2/",
		},
		Storage: &repository.Storage{
			BlobStoreName: "default",
		},
	}
}
