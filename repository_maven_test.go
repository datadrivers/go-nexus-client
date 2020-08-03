package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryMavenGroupRead(t *testing.T) {
	client := getTestClient()

	repoName := "maven-public"

	repo, err := client.RepositoryRead(repoName)
	assert.Nil(t, err)
	assert.NotNil(t, repo)
	assert.Equal(t, repoName, repo.Name)
	assert.Equal(t, RepositoryFormatMaven2, repo.Format)
	assert.Equal(t, RepositoryTypeGroup, repo.Type)
	assert.NotNil(t, repo.RepositoryGroup)
	assert.Greater(t, len(repo.RepositoryGroup.MemberNames), 0)
	assert.NotNil(t, repo.RepositoryStorage)
	assert.Equal(t, "default", repo.RepositoryStorage.BlobStoreName)
}

func TestRepositoryMavenHosted(t *testing.T) {
	client := getTestClient()

	repo := getTestRepositoryMavenHosted("test-maven-repo-hosted", "STRICT", "RELEASE")

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	createdRepo, err := client.RepositoryRead(repo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdRepo)

	if createdRepo != nil {
		writePolicy := "ALLOW"
		createdRepo.RepositoryMaven.LayoutPolicy = "PERMISSIVE"
		createdRepo.RepositoryStorage.WritePolicy = &writePolicy
		err := client.RepositoryUpdate(createdRepo.Name, *createdRepo)
		assert.Nil(t, err)

		updatedRepo, err := client.RepositoryRead(createdRepo.Name)
		assert.Nil(t, err)
		assert.NotNil(t, updatedRepo)
		assert.Equal(t, updatedRepo.RepositoryMaven.LayoutPolicy, "PERMISSIVE")
		assert.Equal(t, *updatedRepo.RepositoryStorage.WritePolicy, "ALLOW")

		err = client.RepositoryDelete(createdRepo.Name)
		assert.Nil(t, err)
	}
}

func getTestRepositoryMavenHosted(name, layoutPolicy, versionPoliy string) Repository {
	writePolicy := "ALLOW_ONCE"
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
			WritePolicy:   &writePolicy,
		},
	}
}
