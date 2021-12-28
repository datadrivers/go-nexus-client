package legacy

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func TestLegacyRepositoryBowerHosted(t *testing.T) {
	service := getTestService()

	// Create hosted bower repo
	repo := getTestLegacyRepositoryBowerHosted("test-legacy-bower-hosted-" + strconv.Itoa(rand.Intn(1024)))
	err := service.Create(repo)
	assert.Nil(t, err)

	proxyRepo := getTestLegacyRepositoryBowerProxy("test-legacy-bower-proxy-" + strconv.Itoa(rand.Intn(1024)))
	err = service.Create(proxyRepo)
	assert.Nil(t, err)

	// Create bower group repo
	groupRepo := getTestLegacyRepositoryBowerGroup("test-legacy-bower-group-"+strconv.Itoa(rand.Intn(1024)), []string{repo.Name, proxyRepo.Name})
	err = service.Create(groupRepo)
	assert.Nil(t, err)

	updatedGroupRepo := groupRepo
	updatedGroupRepo.Online = false

	err = service.Update(groupRepo.Name, updatedGroupRepo)
	assert.Nil(t, err)

	err = service.Delete(groupRepo.Name)
	assert.Nil(t, err)

	err = service.Delete(proxyRepo.Name)
	assert.Nil(t, err)

	err = service.Delete(repo.Name)
	assert.Nil(t, err)
}

func getTestLegacyRepositoryBowerHosted(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Type:   repository.RepositoryTypeHosted,
		Format: repository.RepositoryFormatBower,
		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
			WritePolicy:   (*repository.StorageWritePolicy)(tools.GetStringPointer("ALLOW_ONCE")),
		},
	}
}

func getTestLegacyRepositoryBowerGroup(name string, memberNames []string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Format: repository.RepositoryFormatBower,
		Type:   repository.RepositoryTypeGroup,
		Online: true,
		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
		},
		Group: &repository.Group{
			MemberNames: memberNames,
		},
	}
}

func getTestLegacyRepositoryBowerProxy(name string) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Format: repository.RepositoryFormatBower,
		Type:   repository.RepositoryTypeProxy,
		Bower: &repository.Bower{
			RewritePackageUrls: true,
		},
		Cleanup: &repository.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		HTTPClient: &repository.HTTPClient{
			Connection: &repository.HTTPClientConnection{
				Timeout: tools.GetIntPointer(20),
			},
		},
		NegativeCache: &repository.NegativeCache{
			Enabled: true,
		},
		Proxy: &repository.Proxy{
			RemoteURL: tools.GetStringPointer("https://registry.bower.io"),
		},
		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
		},
	}
}
