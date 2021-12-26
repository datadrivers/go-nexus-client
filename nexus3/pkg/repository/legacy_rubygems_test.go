package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func TestLegacyRepositoryRubyHosted(t *testing.T) {
	service := getTestService()

	testHostedRepo := getTestLegacyRepositoryRubyHosted("test-repo-ruby-hosted")
	hostedCreateErr := service.Legacy.Create(testHostedRepo)
	assert.Nil(t, hostedCreateErr)
	hostedRead, hostedReadErr := service.Legacy.Get(testHostedRepo.Name)
	assert.Nil(t, hostedReadErr)
	assert.Equal(t, testHostedRepo.Type, hostedRead.Type)

	testProxyRepo := getTestLegacyRepositoryRubyProxy("test-repo-ruby-proxy")
	proxyCreateErr := service.Legacy.Create(testProxyRepo)
	assert.Nil(t, proxyCreateErr)
	proxyRead, proxyReadErr := service.Legacy.Get(testProxyRepo.Name)
	assert.Nil(t, proxyReadErr)
	assert.Equal(t, testProxyRepo.Type, proxyRead.Type)

	testGroupRepo := getTestLegacyRepositoryRubyGroup("test-repo-ruby-group", []string{testHostedRepo.Name, testProxyRepo.Name})
	groupCreateErr := service.Legacy.Create(testGroupRepo)
	assert.Nil(t, groupCreateErr)
	groupRead, groupReadErr := service.Legacy.Get(testGroupRepo.Name)
	assert.Nil(t, groupReadErr)
	assert.Equal(t, testGroupRepo.Type, groupRead.Type)
	assert.ElementsMatch(t, testGroupRepo.MemberNames, groupRead.MemberNames)

	_ = service.Legacy.Delete(testGroupRepo.Name)
	_ = service.Legacy.Delete(testProxyRepo.Name)
	_ = service.Legacy.Delete(testHostedRepo.Name)
}

func getTestLegacyRepositoryRubyHosted(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Format: repository.RepositoryFormatRuby,
		Type:   repository.RepositoryTypeHosted,
		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
			WritePolicy:   (*repository.StorageWritePolicy)(tools.GetStringPointer("ALLOW_ONCE")),
		},
		Cleanup: &repository.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
	}
}

func getTestLegacyRepositoryRubyProxy(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:       name,
		Format:     repository.RepositoryFormatRuby,
		Type:       repository.RepositoryTypeProxy,
		HTTPClient: &repository.HTTPClient{},
		NegativeCache: &repository.NegativeCache{
			Enabled: true,
		},
		Proxy: &repository.Proxy{
			RemoteURL: tools.GetStringPointer("https://rubygems.org/"),
		},
		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
		},
	}
}

func getTestLegacyRepositoryRubyGroup(name string, memberNames []string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Format: repository.RepositoryFormatRuby,
		Type:   repository.RepositoryTypeGroup,
		Group: &repository.Group{
			MemberNames: memberNames,
		},
		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
		},
	}
}
