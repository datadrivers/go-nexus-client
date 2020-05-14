package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		createdRepo.RepositoryStorage.WritePolicy = "ALLOW"
		err := client.RepositoryUpdate(createdRepo.Name, *createdRepo)
		assert.Nil(t, err)

		updatedRepo, err := client.RepositoryRead(createdRepo.Name)
		assert.Nil(t, err)
		assert.NotNil(t, updatedRepo)
		assert.Equal(t, updatedRepo.RepositoryMaven.LayoutPolicy, "PERMISSIVE")
		assert.Equal(t, updatedRepo.RepositoryStorage.WritePolicy, "ALLOW")

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
