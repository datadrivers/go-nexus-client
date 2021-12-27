package maven

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestMavenProxyRepository(name string) repository.MavenProxyRepository {
	versionPolicy := repository.MavenVersionPolicyRelease
	layoutPolicy := repository.MavenLayoutPolicyStrict
	return repository.MavenProxyRepository{
		Name:   name,
		Online: true,
		HTTPClient: repository.HTTPClientWithPreemptiveAuth{
			Blocked:   true,
			AutoBlock: true,
			Authentication: &repository.HTTPClientAuthenticationWithPreemptive{
				Type:       repository.HTTPClientAuthenticationTypeUsername,
				Username:   tools.GetStringPointer("user"),
				Password:   tools.GetStringPointer("password"),
				Preemptive: tools.GetBoolPointer(false),
			},
		},
		NegativeCache: repository.NegativeCache{
			Enabled: true,
			TTL:     1440,
		},
		Proxy: repository.Proxy{
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
			RemoteURL:      tools.GetStringPointer("https://archive.ubuntu.com/ubuntu/"),
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Maven: repository.Maven{
			VersionPolicy: &versionPolicy,
			LayoutPolicy:  &layoutPolicy,
		},
	}
}

func TestMavenProxyRepository(t *testing.T) {
	service := getTestService()
	repo := getTestMavenProxyRepository("test-maven-repo-hosted")

	err := service.Proxy.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Proxy.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.HTTPClient.Blocked, generatedRepo.HTTPClient.Blocked)
	assert.Equal(t, repo.HTTPClient.AutoBlock, generatedRepo.HTTPClient.AutoBlock)
	assert.Equal(t, repo.HTTPClient.Authentication.Type, generatedRepo.HTTPClient.Authentication.Type)
	assert.Equal(t, repo.HTTPClient.Authentication.Preemptive, generatedRepo.HTTPClient.Authentication.Preemptive)
	assert.Equal(t, repo.HTTPClient.Authentication.Username, generatedRepo.HTTPClient.Authentication.Username)
	assert.Equal(t, repo.NegativeCache, generatedRepo.NegativeCache)
	assert.Equal(t, repo.Proxy, generatedRepo.Proxy)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)
	assert.Equal(t, repo.Maven, generatedRepo.Maven)

	newLayoutPolicy := repository.MavenLayoutPolicyPermissive
	updatedRepo := repo
	updatedRepo.Online = false
	updatedRepo.HTTPClient.Authentication.Preemptive = tools.GetBoolPointer(true)
	updatedRepo.LayoutPolicy = &newLayoutPolicy

	err = service.Proxy.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Proxy.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	assert.Equal(t, updatedRepo.HTTPClient.Authentication.Preemptive, generatedRepo.HTTPClient.Authentication.Preemptive)
	assert.Equal(t, updatedRepo.LayoutPolicy, generatedRepo.LayoutPolicy)

	service.Proxy.Delete(repo.Name)
	assert.Nil(t, err)
}
