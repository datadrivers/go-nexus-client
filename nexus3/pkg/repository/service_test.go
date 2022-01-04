package repository

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

func getTestService() *RepositoryService {
	return NewRepositoryService(getTestClient())
}

func getDefaultConfig() client.Config {
	return client.Config{
		Insecure: tools.GetEnv("NEXUS_INSECURE_SKIP_VERIFY", true).(bool),
		Password: tools.GetEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      tools.GetEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: tools.GetEnv("NEXUS_USRNAME", "admin").(string),
	}
}

func TestNewRepositoryService(t *testing.T) {
	s := getTestService()

	assert.NotNil(t, s, "NewRepositoryService() must not return nil")
}

func TestListRepositories(t *testing.T) {
	service := getTestService()
	assert.NotNil(t, service, "NewRepositoryService() must not return nil")

	repoInfos, err := service.List()
	assert.Nil(t, err)
	assert.NotEmpty(t, repoInfos)
	assert.NotEmpty(t, repoInfos[0].Name)
	assert.NotEmpty(t, repoInfos[0].Format)
	assert.NotEmpty(t, repoInfos[0].Type)
	assert.NotEmpty(t, repoInfos[0].URL)
}
