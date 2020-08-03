package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryRubyHosted(t *testing.T) {
	client := getTestClient()

	testHostedRepo := getTestRepositoryRubyHosted("test-repo-ruby-hosted")
	hostedCreateErr := client.RepositoryCreate(testHostedRepo)
	assert.Nil(t, hostedCreateErr)
	hostedRead, hostedReadErr := client.RepositoryRead(testHostedRepo.Name)
	assert.Nil(t, hostedReadErr)
	assert.Equal(t, testHostedRepo.Type, hostedRead.Type)

	testProxyRepo := getTestRepositoryRubyProxy("test-repo-ruby-proxy")
	proxyCreateErr := client.RepositoryCreate(testProxyRepo)
	assert.Nil(t, proxyCreateErr)
	proxyRead, proxyReadErr := client.RepositoryRead(testProxyRepo.Name)
	assert.Nil(t, proxyReadErr)
	assert.Equal(t, testProxyRepo.Type, proxyRead.Type)

	testGroupRepo := getTestRepositoryRubyGroup("test-repo-ruby-group", []string{testHostedRepo.Name, testProxyRepo.Name})
	groupCreateErr := client.RepositoryCreate(testGroupRepo)
	assert.Nil(t, groupCreateErr)
	groupRead, groupReadErr := client.RepositoryRead(testGroupRepo.Name)
	assert.Nil(t, groupReadErr)
	assert.Equal(t, testGroupRepo.Type, groupRead.Type)
	assert.ElementsMatch(t, testGroupRepo.MemberNames, groupRead.MemberNames)

	_ = client.RepositoryDelete(testGroupRepo.Name)
	_ = client.RepositoryDelete(testProxyRepo.Name)
	_ = client.RepositoryDelete(testHostedRepo.Name)
}

func getTestRepositoryRubyHosted(name string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatRuby,
		Type:   RepositoryTypeHosted,
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
			WritePolicy:   makeStringAddressable("ALLOW_ONCE"),
		},
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
	}
}

func getTestRepositoryRubyProxy(name string) Repository {
	return Repository{
		Name:                 name,
		Format:               RepositoryFormatRuby,
		Type:                 RepositoryTypeProxy,
		RepositoryHTTPClient: &RepositoryHTTPClient{},
		RepositoryNegativeCache: &RepositoryNegativeCache{
			Enabled: true,
		},
		RepositoryProxy: &RepositoryProxy{
			RemoteURL: "https://rubygems.org/",
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
		},
	}
}

func getTestRepositoryRubyGroup(name string, memberNames []string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatRuby,
		Type:   RepositoryTypeGroup,
		RepositoryGroup: &RepositoryGroup{
			MemberNames: memberNames,
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
		},
	}
}
