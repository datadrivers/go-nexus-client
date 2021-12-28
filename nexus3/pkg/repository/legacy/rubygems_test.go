package legacy

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func TestLegacyRepositoryRubyHosted(t *testing.T) {
	service := getTestService()

	testHostedRepo := getTestLegacyRepositoryRubyHosted("test-legacy-ruby-hosted-" + strconv.Itoa(rand.Intn(1024)))
	hostedCreateErr := service.Create(testHostedRepo)
	assert.Nil(t, hostedCreateErr)
	hostedRead, hostedReadErr := service.Get(testHostedRepo.Name)
	assert.Nil(t, hostedReadErr)
	assert.Equal(t, testHostedRepo.Type, hostedRead.Type)

	testProxyRepo := getTestLegacyRepositoryRubyProxy("test-legacy-ruby-proxy-" + strconv.Itoa(rand.Intn(1024)))
	proxyCreateErr := service.Create(testProxyRepo)
	assert.Nil(t, proxyCreateErr)
	proxyRead, proxyReadErr := service.Get(testProxyRepo.Name)
	assert.Nil(t, proxyReadErr)
	assert.Equal(t, testProxyRepo.Type, proxyRead.Type)

	testGroupRepo := getTestLegacyRepositoryRubyGroup("test-legacy-ruby-group-"+strconv.Itoa(rand.Intn(1024)), []string{testHostedRepo.Name, testProxyRepo.Name})
	groupCreateErr := service.Create(testGroupRepo)
	assert.Nil(t, groupCreateErr)
	groupRead, groupReadErr := service.Get(testGroupRepo.Name)
	assert.Nil(t, groupReadErr)
	assert.Equal(t, testGroupRepo.Type, groupRead.Type)
	assert.ElementsMatch(t, testGroupRepo.MemberNames, groupRead.MemberNames)

	_ = service.Delete(testGroupRepo.Name)
	_ = service.Delete(testProxyRepo.Name)
	_ = service.Delete(testHostedRepo.Name)
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
