package blobstore

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
	"github.com/stretchr/testify/assert"
)

func TestBlobstoreAzure(t *testing.T) {
	if tools.GetEnv("SKIP_AZURE_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus blobstore for Azure tests")
	}
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus blobstore for Azure tests")
	}

	azureAccountName := "terraformprovidernexus"
	azureContainerName := "go-nexus-client"
	bsName := "test-blobstore-azure-" + strconv.Itoa(rand.Intn(1024))
	azureAccountKey := tools.GetEnv("AZURE_STORAGE_ACCOUNT_KEY", "test-key").(string)

	service := getTestService()

	bs := &blobstore.Azure{
		Name: bsName,
		BucketConfiguration: blobstore.AzureBucketConfiguration{
			AccountName: azureAccountName,
			Authentication: blobstore.AzureBucketConfigurationAuthentication{
				AuthenticationMethod: blobstore.AzureAuthenticationMethodAccountKey,
				AccountKey:           azureAccountKey,
			},
			ContainerName: azureContainerName,
		},
	}

	err := service.Azure.Create(bs)
	assert.Nil(t, err)

	azureBS, err := service.Azure.Get(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, azureBS)
	assert.Equal(t, bsName, azureBS.Name)
	assert.NotNil(t, azureBS.BucketConfiguration)
	assert.Equal(t, azureAccountName, azureBS.BucketConfiguration.AccountName)
	assert.Equal(t, azureContainerName, azureBS.BucketConfiguration.ContainerName)

	azureBS.SoftQuota = &blobstore.SoftQuota{
		Type:  "spaceRemainingQuota",
		Limit: 100000000,
	}

	err = service.Azure.Update(azureBS.Name, azureBS)
	assert.Nil(t, err)

	updatedBlobstore, err := service.Azure.Get(azureBS.Name)
	assert.Nil(t, err)
	assert.NotNil(t, updatedBlobstore)
	assert.NotNil(t, updatedBlobstore.SoftQuota)

	err = service.Azure.Delete(azureBS.Name)
	assert.Nil(t, err)

}

func TestBlobstoreAzureTestConnection(t *testing.T) {
	if tools.GetEnv("SKIP_AZURE_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus blobstore for Azure tests")
	}

	azureAccountName := "terraformprovidernexus"
	azureContainerName := "go-nexus-client"
	bsName := "test-blobstore-azure"
	azureAccountKey := "test-key"

	service := getTestService()

	bs := &blobstore.Azure{
		Name: bsName,
		BucketConfiguration: blobstore.AzureBucketConfiguration{
			AccountName: azureAccountName,
			Authentication: blobstore.AzureBucketConfigurationAuthentication{
				AuthenticationMethod: blobstore.AzureAuthenticationMethodAccountKey,
				AccountKey:           azureAccountKey,
			},
			ContainerName: azureContainerName,
		},
	}

	err := service.Azure.TestConnection(bs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "could not connect to azure storage account container: Azure Blob Store connection failed")

	if value, exists := os.LookupEnv("AZURE_STORAGE_ACCOUNT_KEY"); exists {
		bs.BucketConfiguration.Authentication.AccountKey = value
		err := service.Azure.TestConnection(bs)
		assert.Nil(t, err)
	}
}
