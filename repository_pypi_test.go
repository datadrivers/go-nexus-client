package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryPyPiHosted(t *testing.T) {
	client := getTestClient()

	hostedRepo := getTestRepositoryPyPiHosted("test-repo-pypi-hosted")
	err := client.RepositoryCreate(hostedRepo)
	assert.Nil(t, err)

	if err == nil {
		proxyRepo := getTestRepositoryPyPiProxy("test-repo-pypi-proxy")
		err = client.RepositoryCreate(proxyRepo)
		assert.Nil(t, err)

		if err == nil {
			groupRepo := getTestRepositoryPyPiGroup("test-repo-pypi-group", []string{hostedRepo.Name, proxyRepo.Name})
			err = client.RepositoryCreate(groupRepo)
			assert.Nil(t, err)

			if err == nil {
				err = client.RepositoryDelete(groupRepo.Name)
				assert.Nil(t, err)
			}

			err = client.RepositoryDelete(proxyRepo.Name)
			assert.Nil(t, err)
		}

		err = client.RepositoryDelete(hostedRepo.Name)
		assert.Nil(t, err)
	}
}

func getTestRepositoryPyPiHosted(name string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatPyPi,
		Type:   RepositoryTypeHosted,
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
	}
}

func getTestRepositoryPyPiProxy(name string) Repository {
	return Repository{
		Name:                 name,
		Format:               RepositoryFormatPyPi,
		Type:                 RepositoryTypeProxy,
		RepositoryHTTPClient: &RepositoryHTTPClient{},
		RepositoryNegativeCache: &RepositoryNegativeCache{
			Enabled: true,
		},
		RepositoryProxy: &RepositoryProxy{
			RemoteURL: "https://pypi.org/",
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
	}
}

func getTestRepositoryPyPiGroup(name string, memberNames []string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatPyPi,
		Type:   RepositoryTypeGroup,
		RepositoryGroup: &RepositoryGroup{
			MemberNames: memberNames,
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
	}
}
