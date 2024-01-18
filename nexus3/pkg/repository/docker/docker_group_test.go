package docker

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
	"github.com/stretchr/testify/assert"
)

func getTestDockerGroupRepository(name string) repository.DockerGroupRepository {
	return repository.DockerGroupRepository{
		Name:   name,
		Online: true,

		Group: repository.GroupDeploy{
			MemberNames: []string{},
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		Docker: repository.Docker{
			ForceBasicAuth: true,
			V1Enabled:      false,
			HTTPPort:       tools.GetIntPointer(8080),
			HTTPSPort:      tools.GetIntPointer(8443),
			Subdomain:      tools.GetStringPointer(name),
		},
	}
}

func TestDockerGroupRepository(t *testing.T) {
	service := getTestService()
	repo := getTestDockerGroupRepository("test-docker-repo-group-" + strconv.Itoa(rand.Intn(1024)))

	testProxyRepo := getTestDockerProxyRepository("test-docker-group-proxy-" + strconv.Itoa(rand.Intn(1024)))
	defer service.Proxy.Delete(testProxyRepo.Name)
	err := service.Proxy.Create(testProxyRepo)
	assert.Nil(t, err)
	repo.Group.MemberNames = append(repo.Group.MemberNames, testProxyRepo.Name)

	err = service.Group.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Group.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Group, generatedRepo.Group)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)
	assert.Equal(t, repo.Docker, generatedRepo.Docker)

	updatedRepo := repo
	updatedRepo.Online = false
	updatedRepo.ForceBasicAuth = false

	err = service.Group.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Group.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	assert.Equal(t, updatedRepo.ForceBasicAuth, generatedRepo.ForceBasicAuth)

	service.Group.Delete(repo.Name)
	assert.Nil(t, err)
}
