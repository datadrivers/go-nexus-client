package npm

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/tools"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

func getTestNpmGroupRepository(name string) repository.NpmGroupRepository {
	return repository.NpmGroupRepository{
		Name:   name,
		Online: true,

		Group: repository.GroupDeploy{
			MemberNames: []string{},
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
	}
}

func TestNpmGroupRepository(t *testing.T) {
	service := getTestService()
	repo := getTestNpmGroupRepository("test-npm-repo-group-" + strconv.Itoa(rand.Intn(1024)))

	testProxyRepo := getTestNpmProxyRepository("test-npm-group-proxy-" + strconv.Itoa(rand.Intn(1024)))
	defer service.Proxy.Delete(testProxyRepo.Name)
	err := service.Proxy.Create(testProxyRepo)
	assert.Nil(t, err)

	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "false" {
		testHostedRepo := getTestNpmHostedRepository("test-npm-group-hosted-" + strconv.Itoa(rand.Intn(1024)))
		defer service.Hosted.Delete(testHostedRepo.Name)
		err = service.Hosted.Create(testHostedRepo)
		assert.Nil(t, err)

		repo.Group.MemberNames = append(repo.Group.MemberNames, testProxyRepo.Name, testHostedRepo.Name)
		repo.Group.WritableMember = &testHostedRepo.Name
	} else {
		repo.Group.MemberNames = append(repo.Group.MemberNames, testProxyRepo.Name)
	}

	err = service.Group.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Group.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Group, generatedRepo.Group)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)

	updatedRepo := repo
	updatedRepo.Online = false

	err = service.Group.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Group.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)

	service.Group.Delete(repo.Name)
	assert.Nil(t, err)
}
