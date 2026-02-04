package terraform

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

// getTestTerraformProxyRepository returns a Terraform proxy repository configuration
// that is suitable for Create/Get/Update/Delete integration tests.
func getTestTerraformProxyRepository(name string) repository.TerraformProxyRepository {
	return repository.TerraformProxyRepository{
		Name:   name,
		Online: true,
		HTTPClient: repository.HTTPClient{
			Blocked:   true,
			AutoBlock: true,
			Connection: &repository.HTTPClientConnection{
				Timeout:       tools.GetIntPointer(20),
				UseTrustStore: tools.GetBoolPointer(true),
			},
		},
		NegativeCache: repository.NegativeCache{
			Enabled: true,
			TTL:     1440,
		},
		Proxy: repository.Proxy{
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
			RemoteURL:      "https://registry.terraform.io",
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		// Cleanup is optional; enable if you want to validate cleanup policy behavior.
		// Cleanup: &repository.Cleanup{
		// 	PolicyNames: []string{"cleanup-policy-name"},
		// },
	}
}

func TestTerraformProxyRepository(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}

	service := getTestService()
	repo := getTestTerraformProxyRepository("test-terraform-repo-proxy-" + strconv.Itoa(rand.Intn(1024)))

	// Create
	err := service.Proxy.Create(repo)
	assert.Nil(t, err)

	// Get and verify
	generatedRepo, err := service.Proxy.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)

	// Validate common repository attributes
	assert.Equal(t, repo.HTTPClient.Blocked, generatedRepo.HTTPClient.Blocked)
	assert.Equal(t, repo.HTTPClient.AutoBlock, generatedRepo.HTTPClient.AutoBlock)
	assert.Equal(t, repo.HTTPClient.Connection.Timeout, generatedRepo.HTTPClient.Connection.Timeout)
	assert.Equal(t, repo.HTTPClient.Connection.UseTrustStore, generatedRepo.HTTPClient.Connection.UseTrustStore)

	assert.Equal(t, repo.NegativeCache, generatedRepo.NegativeCache)
	assert.Equal(t, repo.Proxy, generatedRepo.Proxy)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)

	// Cleanup is optional; only validate if you set it in the test repo.
	if repo.Cleanup != nil {
		assert.NotNil(t, generatedRepo.Cleanup)
		assert.Equal(t, repo.Cleanup, generatedRepo.Cleanup)
	}

	// Update and verify
	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Proxy.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)

	generatedRepo, err = service.Proxy.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)

	// Delete
	err = service.Proxy.Delete(repo.Name)
	assert.Nil(t, err)
}
