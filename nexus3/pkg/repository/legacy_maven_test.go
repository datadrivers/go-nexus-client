package repository

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func TestLegacyRepositoryMavenGroupRead(t *testing.T) {
	service := getTestService()

	repoName := "maven-public"

	repo, err := service.Legacy.Get(repoName)
	assert.Nil(t, err)
	assert.NotNil(t, repo)
	assert.Equal(t, repoName, repo.Name)
	assert.Equal(t, repository.RepositoryFormatMaven2, repo.Format)
	assert.Equal(t, repository.RepositoryTypeGroup, repo.Type)
	assert.NotNil(t, repo.Group)
	assert.Greater(t, len(repo.Group.MemberNames), 0)
	assert.NotNil(t, repo.Storage)
	assert.Equal(t, "default", repo.Storage.BlobStoreName)
}

func TestLegacyRepositoryMavenHosted(t *testing.T) {
	service := getTestService()
	layoutPolicy := repository.MavenLayoutPolicyStrict
	versionPolicy := repository.MavenVersionPolicyRelease
	repo := getTestLegacyRepositoryMavenHosted("test-maven-repo-hosted", &layoutPolicy, &versionPolicy)

	err := service.Legacy.Create(repo)
	assert.Nil(t, err)

	createdRepo, err := service.Legacy.Get(repo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdRepo)

	writePolicy := (repository.StorageWritePolicy)("ALLOW")
	layoutPolicy = repository.MavenLayoutPolicyPermissive
	createdRepo.Maven.LayoutPolicy = &layoutPolicy
	createdRepo.Storage.WritePolicy = &writePolicy
	err = service.Legacy.Update(createdRepo.Name, *createdRepo)
	assert.Nil(t, err)

	updatedRepo, err := service.Legacy.Get(createdRepo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, updatedRepo)
	assert.Equal(t, repository.MavenLayoutPolicyPermissive, *updatedRepo.Maven.LayoutPolicy)
	assert.Equal(t, repository.StorageWritePolicyAllow, *updatedRepo.Storage.WritePolicy)

	err = service.Legacy.Delete(createdRepo.Name)
	assert.Nil(t, err)

}

func getTestLegacyRepositoryMavenHosted(name string, layoutPolicy *repository.MavenLayoutPolicy, versionPolicy *repository.MavenVersionPolicy) repository.LegacyRepository {
	return repository.LegacyRepository{
		Name:   name,
		Format: repository.RepositoryFormatMaven2,
		Type:   repository.RepositoryTypeHosted,
		Online: true,

		Maven: &repository.Maven{
			LayoutPolicy:  layoutPolicy,
			VersionPolicy: versionPolicy,
		},

		Storage: &repository.HostedStorage{
			BlobStoreName: "default",
			WritePolicy:   (*repository.StorageWritePolicy)(tools.GetStringPointer("ALLOW_ONCE")),
		},
	}
}
