package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryYumHosted(t *testing.T) {
	client := NewClient(getDefaultConfig())

	repoName := "tst-yum-repo-hosted"

	repo := getTestRepositoryYumHosted(repoName)

	err := client.RepositoryCreate(repo)
	assert.Nil(t, err)

	if err == nil {
		createdRepo, err := client.RepositoryRead(repo.Name)
		assert.Nil(t, err)
		assert.NotNil(t, createdRepo)

		assert.Equal(t, repo.Name, createdRepo.Name)
		assert.Equal(t, repo.Type, createdRepo.Type)
		assert.Equal(t, repo.Format, createdRepo.Format)
		assert.Equal(t, repo.Online, createdRepo.Online)

		assert.Equal(t, repo.RepositoryYum.DeployPolicy, createdRepo.RepositoryYum.DeployPolicy)
		assert.Equal(t, repo.RepositoryYum.RepodataDepth, createdRepo.RepositoryYum.RepodataDepth)

		createdRepo.RepositoryYum.DeployPolicy = "PERMISSIVE"
		err = client.RepositoryUpdate(createdRepo.Name, *createdRepo)
		assert.Nil(t, err)

		err = client.RepositoryDelete(repo.Name)
		assert.Nil(t, err)
	}
}

func getTestRepositoryYumHosted(name string) Repository {
	return Repository{
		Name:   name,
		Format: RepositoryFormatYum,
		Type:   RepositoryTypeHosted,
		Online: true,

		RepositoryStorage: &RepositoryStorage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},

		RepositoryYum: &RepositoryYum{
			DeployPolicy:  "STRICT",
			RepodataDepth: 1,
		},
	}
}
