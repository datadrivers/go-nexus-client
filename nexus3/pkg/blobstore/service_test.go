package blobstore

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/stretchr/testify/assert"
)

var (
	testClient *client.Client = nil
)

func getTestClient() *client.Client {
	if testClient != nil {
		return testClient
	}
	return client.NewClient(getDefaultConfig())
}

func getTestService() *BlobStoreService {
	return NewBlobStoreService(getTestClient())
}

func getDefaultConfig() client.Config {
	return client.Config{
		Insecure: tools.GetEnv("NEXUS_INSECURE_SKIP_VERIFY", true).(bool),
		Password: tools.GetEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      tools.GetEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: tools.GetEnv("NEXUS_USRNAME", "admin").(string),
	}
}

func TestNewBlobStoreService(t *testing.T) {
	service := getTestService()

	assert.NotNil(t, service, "NewBlobStoreService() must not return nil")
}

func TestListBlobstores(t *testing.T) {
	service := getTestService()
	blobstores, err := service.List()

	assert.Nil(t, err)
	assert.Equal(t, "default", blobstores[0].Name)
	assert.Equal(t, "File", blobstores[0].Type)
	assert.Equal(t, false, blobstores[0].Unavailable)
	assert.Equal(t, 0, blobstores[0].BlobCount)
}
