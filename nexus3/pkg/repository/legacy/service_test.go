package legacy

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
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

func getTestService() *RepositoryLegacyService {
	return NewRepositoryLegacyService(getTestClient())
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

func TestLegacyRepositoryRead(t *testing.T) {
	service := getTestService()

	name := "maven-central"

	repo, err := service.Get(name)
	assert.Nil(t, err)
	assert.NotNil(t, repo)

	if repo != nil {
		assert.Equal(t, name, repo.Name)
		assert.NotNil(t, repo.Proxy)
	}
}

func TestJSONUnmarshalRepositories(t *testing.T) {
	data := []byte(testJSONUnmarshalRepositories())
	repositories, err := jsonUnmarshalRepositories(data)
	assert.Nil(t, err)
	assert.NotNil(t, repositories)
	assert.Equal(t, 1, len(repositories))

	repo := repositories[0]
	assert.Equal(t, repo.Format, repository.RepositoryFormatMaven2)
	assert.Equal(t, repo.Type, repository.RepositoryTypeProxy)
	assert.Nil(t, repo.Docker)
	assert.NotNil(t, repo.HTTPClient)
}

func testJSONUnmarshalRepositories() string {
	return `[
	{
		"format": "maven2",
		"name": "maven-central",
		"online": true,
		"type": "proxy",
		"cleanup": {
			"policyNames": []
		},
		"docker": null,
		"httpClient": {
			"authentication": {
				"ntlmDomain": "",
				"ntlmHost": "",
				"type": "",
				"username": ""
			},
			"autoBlock": false,
			"blocked": false,
			"connection": {
				"enableCircularRedirects": false,
				"enableCookies": false,
				"retries": 0,
				"timeout": 0,
				"userAgentSuffix": ""
			}
		},
		"negativeCache": {
			"enabled": true,
			"timeToLive": 1440
		},
		"proxy": {
			"contentMaxAge": -1,
			"metadataMaxAge": 1440,
			"remoteUrl": "https://repo1.maven.org/maven2/"
		},
		"storage": {
			"blobStoreName": "default",
			"strictContentTypeValidation": false,
			"writePolicy": "ALLOW"
		}
	}]`
}

func TestLegacyRepositoryFixFormat(t *testing.T) {
	for _, format := range repository.RepositoryFormats {
		if format == repository.RepositoryFormatMaven2 {
			assert.Equal(t, fixRepositoryFormat(repository.RepositoryFormatMaven2), "maven")
		} else {
			assert.Equal(t, fixRepositoryFormat(format), format)
		}
	}
}
