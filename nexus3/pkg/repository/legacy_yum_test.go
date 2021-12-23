package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func TestLegacyRepositoryYumHosted(t *testing.T) {
	service := getTestService()

	repoName := "tst-yum-repo-hosted"

	repo := getTestLegacyRepositoryYumHosted(repoName)

	err := service.Legacy.Create(repo)
	assert.Nil(t, err)

	if err == nil {
		createdRepo, err := service.Legacy.Get(repo.Name)
		assert.Nil(t, err)
		assert.NotNil(t, createdRepo)

		assert.Equal(t, repo.Name, createdRepo.Name)
		assert.Equal(t, repo.Type, createdRepo.Type)
		assert.Equal(t, repo.Format, createdRepo.Format)
		assert.Equal(t, repo.Online, createdRepo.Online)

		assert.Equal(t, repo.Yum.DeployPolicy, createdRepo.Yum.DeployPolicy)
		assert.Equal(t, repo.Yum.RepodataDepth, createdRepo.Yum.RepodataDepth)

		deployPolicy := repository.YumDeployPolicyPermissive
		createdRepo.Yum.DeployPolicy = &deployPolicy
		err = service.Legacy.Update(createdRepo.Name, *createdRepo)
		assert.Nil(t, err)

		err = service.Legacy.Delete(repo.Name)
		assert.Nil(t, err)
	}
}

func getTestLegacyRepositoryYumHosted(name string) repository.LegacyRepository {
	deployPolicy := repository.YumDeployPolicyStrict
	return repository.LegacyRepository{
		Name:   name,
		Format: repository.RepositoryFormatYum,
		Type:   repository.RepositoryTypeHosted,
		Online: true,

		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
			WritePolicy:   tools.GetStringPointer("ALLOW_ONCE"),
		},

		Yum: &repository.Yum{
			DeployPolicy:  &deployPolicy,
			RepodataDepth: 1,
		},
	}
}
