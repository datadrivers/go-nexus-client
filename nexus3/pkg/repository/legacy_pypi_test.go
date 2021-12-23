package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func TestLegacyRepositoryPyPiHosted(t *testing.T) {
	service := getTestService()

	hostedRepo := getTestLegacyRepositoryPyPiHosted("test-repo-pypi-hosted")
	err := service.Legacy.Create(hostedRepo)
	assert.Nil(t, err)

	if err == nil {
		proxyRepo := getTestLegacyRepositoryPyPiProxy("test-repo-pypi-proxy")
		err = service.Legacy.Create(proxyRepo)
		assert.Nil(t, err)

		if err == nil {
			groupRepo := getTestLegacyRepositoryPyPiGroup("test-repo-pypi-group", []string{hostedRepo.Name, proxyRepo.Name})
			err = service.Legacy.Create(groupRepo)
			assert.Nil(t, err)

			if err == nil {
				err = service.Legacy.Delete(groupRepo.Name)
				assert.Nil(t, err)
			}

			err = service.Legacy.Delete(proxyRepo.Name)
			assert.Nil(t, err)
		}

		err = service.Legacy.Delete(hostedRepo.Name)
		assert.Nil(t, err)
	}
}

func getTestLegacyRepositoryPyPiHosted(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Format: repository.RepositoryFormatPyPi,
		Type:   repository.RepositoryTypeHosted,
		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
			WritePolicy:   tools.GetStringPointer("ALLOW_ONCE"),
		},
		Cleanup: &repository.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
	}
}

func getTestLegacyRepositoryPyPiProxy(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Format: repository.RepositoryFormatPyPi,
		Type:   repository.RepositoryTypeProxy,
		HTTPClient: &repository.HTTPClient{
			Connection: &repository.HTTPClientConnection{
				Timeout: tools.GetIntPointer(20),
			},
		},
		NegativeCache: &repository.NegativeCache{
			Enabled: true,
		},
		Proxy: &repository.Proxy{
			RemoteURL: tools.GetStringPointer("https://pypi.org/"),
		},
		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
			WritePolicy:   tools.GetStringPointer("ALLOW_ONCE"),
		},
	}
}

func getTestLegacyRepositoryPyPiGroup(name string, memberNames []string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Format: repository.RepositoryFormatPyPi,
		Type:   repository.RepositoryTypeGroup,
		Group: &repository.Group{
			MemberNames: memberNames,
		},
		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
			WritePolicy:   tools.GetStringPointer("ALLOW_ONCE"),
		},
	}
}
