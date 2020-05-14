package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryRead(t *testing.T) {
	client := NewClient(getDefaultConfig())

	name := "maven-central"

	repo, err := client.RepositoryRead(name)
	assert.Nil(t, err)
	assert.NotNil(t, repo)

	if repo != nil {
		assert.Equal(t, name, repo.Name)
		assert.NotNil(t, repo.RepositoryProxy)
	}
}

func TestJSONUnmarshalRepositories(t *testing.T) {
	data := []byte(testJSONUnmarshalRepositories())
	repositories, err := jsonUnmarshalRepositories(data)
	assert.Nil(t, err)
	assert.NotNil(t, repositories)
	assert.Equal(t, 1, len(repositories))

	repo := repositories[0]
	assert.Equal(t, repo.Format, RepositoryFormatMaven2)
	assert.Equal(t, repo.Type, RepositoryTypeProxy)
	assert.Nil(t, repo.RepositoryDocker)
	assert.NotNil(t, repo.RepositoryHTTPClient)
}

func testJSONUnmarshalRepositories() string {
	return fmt.Sprintf(`[
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
	}]`)
}

func TestRepositoryFixFormat(t *testing.T) {
	for _, format := range RepositoryFormats {
		if format == RepositoryFormatMaven2 {
			assert.Equal(t, fixRepositoryFormat(RepositoryFormatMaven2), "maven")
		} else {
			assert.Equal(t, fixRepositoryFormat(format), format)
		}
	}
}
