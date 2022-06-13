package legacy

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func TestLegacyRepositoryMavenGroupRead(t *testing.T) {
	service := getTestService()

	repoName := "maven-public"

	repo, err := service.Get(repoName)
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
	repo := getTestLegacyRepositoryMavenHosted("test-legacy-maven-hosted-"+strconv.Itoa(rand.Intn(1024)), repository.MavenLayoutPolicyStrict, repository.MavenVersionPolicyRelease)

	err := service.Create(repo)
	assert.Nil(t, err)

	createdRepo, err := service.Get(repo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdRepo)

	writePolicy := (repository.StorageWritePolicy)("ALLOW")
	createdRepo.Maven.LayoutPolicy = repository.MavenLayoutPolicyPermissive
	createdRepo.Storage.WritePolicy = &writePolicy
	err = service.Update(createdRepo.Name, *createdRepo)
	assert.Nil(t, err)

	updatedRepo, err := service.Get(createdRepo.Name)
	assert.Nil(t, err)
	assert.NotNil(t, updatedRepo)
	assert.Equal(t, repository.MavenLayoutPolicyPermissive, updatedRepo.Maven.LayoutPolicy)
	assert.Equal(t, repository.StorageWritePolicyAllow, *updatedRepo.Storage.WritePolicy)

	err = service.Delete(createdRepo.Name)
	assert.Nil(t, err)

}

func getTestLegacyRepositoryMavenHosted(name string, layoutPolicy repository.MavenLayoutPolicy, versionPolicy repository.MavenVersionPolicy) repository.LegacyRepository {
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
