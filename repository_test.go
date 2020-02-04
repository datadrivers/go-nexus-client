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
		assert.NotNil(t, repo.RepositoryCleanup)
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

func getTestRepositoryAptHosted(name string) Repository {
	writePolicy := "ALLOW_ONCE"

	return Repository{
		Name:   name,
		Online: true,
		Type:   RepositoryTypeHosted,
		Format: RepositoryFormatApt,

		RepositoryApt: &RepositoryApt{
			Distribution: "bionic",
		},
		RepositoryAptSigning: &RepositoryAptSigning{
			Keypair:    "string",
			Passphrase: "string",
		},
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 &writePolicy,
		},
	}
}

func TestRepositoryAptHosted(t *testing.T) {
	client := NewClient(getDefaultConfig())
	repo := getTestRepositoryAptHosted("test-apt-repo-hosted")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryUpdate(repo.Name, updatedRepo)
	assert.Nil(t, err)

	err = client.RepositoryDelete(repo.Name)
	assert.Nil(t, err)
}

func getTestRepositoryDockerHostedWithPorts(name string) Repository {
	httpPort := new(int)
	httpsPort := new(int)
	*httpPort = 8082
	*httpsPort = 8083
	writePolicy := "ALLOW_ONCE"

	return Repository{
		Name:   name,
		Online: true,
		Format: RepositoryFormatDocker,
		Type:   RepositoryTypeHosted,
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryDocker: &RepositoryDocker{
			V1Enabled:      false,
			ForceBasicAuth: true,
			HTTPPort:       httpPort,
			HTTPSPort:      httpsPort,
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 &writePolicy,
		},
	}
}

func TestRepositoryDockerHostedWithPorts(t *testing.T) {
	client := NewClient(getDefaultConfig())
	repo := getTestRepositoryDockerHostedWithPorts("test-docker-repo-hosted-with-ports")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryUpdate(repo.Name, updatedRepo)
	assert.Nil(t, err)

	err = client.RepositoryDelete(repo.Name)
	assert.Nil(t, err)
}

func getTestRepositoryDockerHostedWithoutPorts(name string) Repository {
	writePolicy := "ALLOW_ONCE"

	return Repository{
		Name:   name,
		Online: true,
		Format: RepositoryFormatDocker,
		Type:   RepositoryTypeHosted,
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryDocker: &RepositoryDocker{
			V1Enabled:      false,
			ForceBasicAuth: true,
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 &writePolicy,
		},
	}
}

func TestRepositoryDockerHostedWithoutPorts(t *testing.T) {
	client := NewClient(getDefaultConfig())
	repo := getTestRepositoryDockerHostedWithoutPorts("test-docker-repo-hosted-with-ports")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	updatedRepo := repo
	updatedRepo.Online = false

	err = client.RepositoryUpdate(repo.Name, updatedRepo)
	assert.Nil(t, err)

	err = client.RepositoryDelete(repo.Name)
	assert.Nil(t, err)
}

func TestRepositoryMavenRead(t *testing.T) {
	client := NewClient(getDefaultConfig())

	repoName := "maven-public"

	repo, err := client.RepositoryRead(repoName)
	assert.Nil(t, err)
	assert.NotNil(t, repo)
	assert.NotNil(t, repo.RepositoryGroup)
	assert.Greater(t, len(repo.RepositoryGroup.MemberNames), 0)
}

func TestRepositoryDockerProxy(t *testing.T) {
	client := NewClient(getDefaultConfig())
	repo := getTestRepositoryDockerProxy("test-docker-repo-proxy")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	createdRepo, err := client.RepositoryRead(repo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdRepo)

	if createdRepo != nil {
		assert.Equal(t, repo.Name, createdRepo.Name)
		assert.Equal(t, repo.Type, createdRepo.Type)
		assert.Equal(t, repo.Format, createdRepo.Format)

		err := client.RepositoryDelete(repo.Name)
		assert.Nil(t, err)
	}
}

func getTestRepositoryDockerProxy(name string) Repository {
	return Repository{
		Name:   name,
		Online: true,
		Type:   RepositoryTypeProxy,
		Format: RepositoryFormatDocker,
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryDocker: &RepositoryDocker{
			V1Enabled:      false,
			ForceBasicAuth: true,
		},
		RepositoryDockerProxy: &RepositoryDockerProxy{
			IndexType: "HUB",
		},
		RepositoryHTTPClient: &RepositoryHTTPClient{
			Authentication: RepositoryHTTPClientAuthentication{
				Type: "username",
			},
		},
		RepositoryNegativeCache: &RepositoryNegativeCache{},
		RepositoryProxy: &RepositoryProxy{
			RemoteURL: "https://registry-1.docker.io",
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
		},
	}
}
