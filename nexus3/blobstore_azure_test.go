package nexus3

import (
	"os"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/blobstore"
	"github.com/stretchr/testify/assert"
)

func TestBlobstoreAzure(t *testing.T) {
	azureAccountName := "testAcoount"
	azureContainerName := "test-container"
	bsName := "test-blobstore-azure"
	azureAccountKey := getEnv("AZURE_STORAGE_ACCOUNT_KEY", "test-key").(string)

	client := getTestClient()

	bs := &blobstore.Azure{
		Name: bsName,
		BucketConfiguration: blobstore.AzureBucketConfiguration{
			AccountName: azureAccountName,
			Authentication: blobstore.AzureBucketConfigurationAuthentication{
				AuthenticationMethod: BlobstoreAzureAuthenticationMethodAccountKey,
				AccountKey:           azureAccountKey,
			},
			ContainerName: azureContainerName,
		},
	}

	err := client.BlobStore.Azure.Create(bs)
	assert.Nil(t, err)

	azureBS, err := client.BlobStore.Azure.Get(bs.Name)
	assert.Nil(t, err)
	assert.NotNil(t, azureBS)
	if azureBS != nil {
		assert.Equal(t, bsName, azureBS.Name)
		assert.NotNil(t, azureBS.BucketConfiguration)
		assert.Equal(t, azureAccountName, azureBS.BucketConfiguration.AccountName)
		assert.Equal(t, azureContainerName, azureBS.BucketConfiguration.ContainerName)

		azureBS.SoftQuota = &blobstore.SoftQuota{
			Type:  "spaceRemainingQuota",
			Limit: 100000000,
		}

		err = client.BlobStore.Azure.Update(azureBS.Name, azureBS)
		assert.Nil(t, err)

		updatedBlobstore, err := client.BlobStore.Azure.Get(azureBS.Name)
		assert.Nil(t, err)
		assert.NotNil(t, updatedBlobstore)

		if updatedBlobstore != nil {
			assert.NotNil(t, updatedBlobstore.SoftQuota)
		}

		err = client.BlobStore.Azure.Delete(azureBS.Name)
		assert.Nil(t, err)
	}
}

func TestBlobstoreAzureTestConnection(t *testing.T) {
	azureAccountName := "testAcoount"
	azureContainerName := "test-container"
	azureAccountKey := "test-key"
	bsName := "test-blobstore-azure"

	client := getTestClient()

	bs := &blobstore.Azure{
		Name: bsName,
		BucketConfiguration: blobstore.AzureBucketConfiguration{
			AccountName: azureAccountName,
			Authentication: blobstore.AzureBucketConfigurationAuthentication{
				AuthenticationMethod: BlobstoreAzureAuthenticationMethodAccountKey,
				AccountKey:           azureAccountKey,
			},
			ContainerName: azureContainerName,
		},
	}

	err := client.BlobStore.Azure.TestConnection(bs)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "could not connect to azure storage account container: Azure Blob Store connection failed")

	if value, exists := os.LookupEnv("AZURE_STORAGE_ACCOUNT_KEY"); exists {
		bs.BucketConfiguration.Authentication.AccountKey = value
		err := client.BlobStore.Azure.TestConnection(bs)
		assert.Nil(t, err)
	}
}
