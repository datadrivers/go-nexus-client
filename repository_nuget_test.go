package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryNugetProxy(t *testing.T) {
	client := getTestClient()
	repo := getTestRepositoryNugetProxy("test-nuget-proxy-repo")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	createdRepo, err := client.RepositoryRead(repo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdRepo)

	if createdRepo != nil {

		err := client.RepositoryDelete(createdRepo.Name)
		assert.Nil(t, err)
	}
}

func getTestRepositoryNugetProxy(name string) Repository {
	return Repository{
		Format: RepositoryFormatNuget,
		Name:   name,
		Online: true,
		Type:   RepositoryTypeProxy,

		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryHTTPClient:    &RepositoryHTTPClient{
			Connection: &RepositoryHTTPClientConnection{
				Timeout: makeIntAddressable(20),
			},
		},
		RepositoryNegativeCache: &RepositoryNegativeCache{},
		RepositoryNugetProxy: &RepositoryNugetProxy{
			QueryCacheItemMaxAge: 1,
		},
		RepositoryProxy: &RepositoryProxy{
			RemoteURL: "https://www.nuget.org/api/v2/",
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
		},
	}
}
