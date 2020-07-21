package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryBowerHosted(t *testing.T) {
	client := getTestClient()

	// Create hosted bower repo
	repo := getTestRepositoryBowerHosted("test-repo-bower-hosted")
	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	if err == nil {
		proxyRepo := getTestRepositoryBowerProxy("test-repo-bower-proxy")
		err = client.RepositoryCreate(proxyRepo)
		assert.Nil(t, err)

		if err == nil {
			// Create bower group repo
			groupRepo := getTestRepositoryBowerGroup("test-repo-bower-group", []string{repo.Name, proxyRepo.Name})
			err = client.RepositoryCreate(groupRepo)
			assert.Nil(t, err)

			if err == nil {
				updatedGroupRepo := groupRepo
				updatedGroupRepo.Online = false

				err = client.RepositoryUpdate(groupRepo.Name, updatedGroupRepo)
				assert.Nil(t, err)

				err = client.RepositoryDelete(groupRepo.Name)
				assert.Nil(t, err)
			}

			err = client.RepositoryDelete(proxyRepo.Name)
			assert.Nil(t, err)
		}

		err = client.RepositoryDelete(repo.Name)
	}
}

func getTestRepositoryBowerHosted(name string) Repository {
	return Repository{
		Name:   name,
		Type:   RepositoryTypeHosted,
		Format: RepositoryFormatBower,
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
	}
}

func getTestRepositoryBowerGroup(name string, memberNames []string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatBower,
		Type:   RepositoryTypeGroup,
		Online: true,
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
		},
		RepositoryGroup: &RepositoryGroup{
			MemberNames: memberNames,
		},
	}
}

func getTestRepositoryBowerProxy(name string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatBower,
		Type:   RepositoryTypeProxy,
		RepositoryBower: &RepositoryBower{
			RewritePackageUrls: true,
		},
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryHTTPClient: &RepositoryHTTPClient{
			Connection: &RepositoryHTTPClientConnection{
				Timeout: makeIntAddressable(20),
			},
		},
		RepositoryNegativeCache: &RepositoryNegativeCache{
			Enabled: true,
		},
		RepositoryProxy: &RepositoryProxy{
			RemoteURL: "https://registry.bower.io",
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
		},
	}
}
