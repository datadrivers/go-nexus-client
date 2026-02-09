package terraform

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func randSuffix(nBytes int) string {
	b := make([]byte, nBytes)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// getTestTerraformHostedRepository returns a Terraform hosted repository configuration

func getTestTerraformHostedRepository(t *testing.T, name string) (repository.TerraformHostedRepository, bool) {
	t.Helper()

	keypair := os.Getenv("NEXUS_TERRAFORM_SIGNING_KEYPAIR")
	if keypair == "" {
		return repository.TerraformHostedRepository{}, false
	}

	passphrase := os.Getenv("NEXUS_TERRAFORM_SIGNING_PASSPHRASE")
	var passPtr *string
	if passphrase != "" {
		passPtr = &passphrase
	}

	// writePolicy is REQUIRED by HostedStorageAttributes
	writePolicy := repository.StorageWritePolicyAllowOnce // hoặc Allow nếu bạn muốn overwrite

	repo := repository.TerraformHostedRepository{
		Name:   name,
		Online: true,

		Storage: repository.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 &writePolicy,
		},

		TerraformSigning: repository.TerraformSigningAttributes{
			Keypair:    keypair,
			Passphrase: passPtr,
		},
	}

	if os.Getenv("NEXUS_TEST_ENABLE_PROPRIETARY_COMPONENTS") == "true" {
		repo.Component = &repository.Component{
			ProprietaryComponents: true,
		}
	}

	return repo, true
}

func TestTerraformHostedRepository(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}

	service := getTestService()

	repoName := "test-terraform-repo-hosted-" + randSuffix(4)
	repo, ok := getTestTerraformHostedRepository(t, repoName)
	if !ok {
		t.Skip("Missing NEXUS_TERRAFORM_SIGNING_KEYPAIR; skipping Terraform hosted repository test")
	}

	// Create
	err := service.Hosted.Create(repo)
	assert.NoError(t, err)

	// Get and verify
	generatedRepo, err := service.Hosted.Get(repo.Name)
	assert.NoError(t, err)
	assert.NotNil(t, generatedRepo)

	assert.Equal(t, repo.Name, generatedRepo.Name)
	assert.Equal(t, repo.Online, generatedRepo.Online)

	// Storage should round-trip (including writePolicy in most Nexus versions)
	assert.Equal(t, repo.Storage.BlobStoreName, generatedRepo.Storage.BlobStoreName)
	assert.Equal(t, repo.Storage.StrictContentTypeValidation, generatedRepo.Storage.StrictContentTypeValidation)
	if repo.Storage.WritePolicy != nil && generatedRepo.Storage.WritePolicy != nil {
		assert.Equal(t, *repo.Storage.WritePolicy, *generatedRepo.Storage.WritePolicy)
	}

	// Cleanup optional
	if repo.Cleanup != nil {
		assert.NotNil(t, generatedRepo.Cleanup)
		assert.Equal(t, repo.Cleanup, generatedRepo.Cleanup)
	}

	if repo.Component != nil {
		assert.NotNil(t, generatedRepo.Component)
		assert.Equal(t, repo.Component, generatedRepo.Component)
	}

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Hosted.Update(repo.Name, updatedRepo)
	assert.NoError(t, err)

	generatedRepo, err = service.Hosted.Get(repo.Name)
	assert.NoError(t, err)
	assert.NotNil(t, generatedRepo)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)

	// Delete
	err = service.Hosted.Delete(repo.Name)
	assert.NoError(t, err)
}
