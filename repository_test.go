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

func getTestRepositoryAptHosted(name string) Repository {
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
			WritePolicy:                 "ALLOW",
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

func TestRepositoryFormatBower(t *testing.T) {
	client := NewClient(getDefaultConfig())

	// Create hosted bower repo
	hostedRepo := getTestRepositoryBowerHosted("test-repo-bower-hosted")
	err := client.RepositoryCreate(hostedRepo)
	assert.Nil(t, err)

	if err == nil {
		proxyRepo := getTestRepositoryBowerProxy("test-repo-bower-proxy")
		err = client.RepositoryCreate(proxyRepo)
		assert.Nil(t, err)

		if err == nil {
			// Create bower group repo
			groupRepo := getTestRepositoryBowerGroup("test-repo-bower-group", []string{hostedRepo.Name, proxyRepo.Name})
			err = client.RepositoryCreate(groupRepo)
			assert.Nil(t, err)

			if err == nil {
				updatedGroupRepo := groupRepo
				updatedGroupRepo.Online = false

				err = client.RepositoryUpdate(groupRepo.Name, updatedGroupRepo)
				assert.Nil(t, err)

				err = client.RepositoryDelete(groupRepo.Name)
				assert.Nil(t, err)
			}

			err = client.RepositoryDelete(proxyRepo.Name)
			assert.Nil(t, err)
		}

		err = client.RepositoryDelete(hostedRepo.Name)
	}
}

func getTestRepositoryBowerHosted(name string) Repository {
	return Repository{
		Name:   name,
		Type:   RepositoryTypeHosted,
		Format: RepositoryFormatBower,
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
	}
}

func getTestRepositoryBowerGroup(name string, memberNames []string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatBower,
		Type:   RepositoryTypeGroup,
		Online: true,
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
		},
		RepositoryGroup: &RepositoryGroup{
			MemberNames: memberNames,
		},
	}
}

func getTestRepositoryBowerProxy(name string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatBower,
		Type:   RepositoryTypeProxy,
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		RepositoryHTTPClient: &RepositoryHTTPClient{
			Authentication: RepositoryHTTPClientAuthentication{
				Type: "username",
			},
		},
		RepositoryNegativeCache: &RepositoryNegativeCache{
			Enabled: true,
		},
		RepositoryProxy: &RepositoryProxy{
			RemoteURL: "https://registry.bower.io",
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
		},
	}
}

func getTestRepositoryDockerHostedWithPorts(name string) Repository {
	httpPort := new(int)
	httpsPort := new(int)
	*httpPort = 8082
	*httpsPort = 8083
	writePolicy := "ALLOW"

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
			WritePolicy:                 writePolicy,
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
			WritePolicy:                 "ALLOW_ONCE",
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

func TestRepositoryMavenHosted(t *testing.T) {
	client := NewClient(getDefaultConfig())

	repo := getTestRepositoryMavenHosted("test-maven-repo-hosted", "STRICT", "RELEASE")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	createdRepo, err := client.RepositoryRead(repo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdRepo)

	if createdRepo != nil {

		createdRepo.RepositoryMaven.LayoutPolicy = "PERMISSIVE"
		err := client.RepositoryUpdate(createdRepo.Name, *createdRepo)
		assert.Nil(t, err)

		err = client.RepositoryDelete(createdRepo.Name)
		assert.Nil(t, err)
	}
}

func getTestRepositoryMavenHosted(name, layoutPolicy, versionPoliy string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatMaven2,
		Type:   RepositoryTypeHosted,
		Online: true,

		RepositoryMaven: &RepositoryMaven{
			LayoutPolicy:  layoutPolicy,
			VersionPolicy: versionPoliy,
		},

		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
	}
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

func TestRepositoryFormatPyPi(t *testing.T) {
	client := NewClient(getDefaultConfig())

	hostedRepo := getTestRepositoryPyPiHosted("test-repo-pypi-hosted")
	err := client.RepositoryCreate(hostedRepo)
	assert.Nil(t, err)

	if err == nil {
		proxyRepo := getTestRepositoryPyPiProxy("test-repo-pypi-proxy")
		err = client.RepositoryCreate(proxyRepo)
		assert.Nil(t, err)

		if err == nil {
			groupRepo := getTestRepositoryPyPiGroup("test-repo-pypi-group", []string{hostedRepo.Name, proxyRepo.Name})
			err = client.RepositoryCreate(groupRepo)
			assert.Nil(t, err)

			if err == nil {
				err = client.RepositoryDelete(groupRepo.Name)
				assert.Nil(t, err)
			}

			err = client.RepositoryDelete(proxyRepo.Name)
			assert.Nil(t, err)
		}

		err = client.RepositoryDelete(hostedRepo.Name)
		assert.Nil(t, err)
	}
}

func getTestRepositoryPyPiHosted(name string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatPyPi,
		Type:   RepositoryTypeHosted,
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
		RepositoryCleanup: &RepositoryCleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
	}
}

func getTestRepositoryPyPiProxy(name string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatPyPi,
		Type:   RepositoryTypeProxy,
		RepositoryHTTPClient: &RepositoryHTTPClient{
			Authentication: RepositoryHTTPClientAuthentication{
				Type: "username",
			},
		},
		RepositoryNegativeCache: &RepositoryNegativeCache{
			Enabled: true,
		},
		RepositoryProxy: &RepositoryProxy{
			RemoteURL: "https://pypi.org/",
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
	}
}

func getTestRepositoryPyPiGroup(name string, memberNames []string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatPyPi,
		Type:   RepositoryTypeGroup,
		RepositoryGroup: &RepositoryGroup{
			MemberNames: memberNames,
		},
		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
	}
}
