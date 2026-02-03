package apt_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/tools"
	"github.com/williamt1997/go-nexus-client/nexus3/schema"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

func getTestAptProxyRepository(name string) repository.AptProxyRepository {
	return repository.AptProxyRepository{
		Name:   name,
		Online: true,
		Apt: repository.AptProxy{
			Distribution: "bionic",
			Flat:         true,
		},
		HTTPClient: repository.HTTPClient{
			Blocked:   true,
			AutoBlock: true,
			Connection: &repository.HTTPClientConnection{
				Timeout:       tools.GetIntPointer(20),
				UseTrustStore: tools.GetBoolPointer(true),
			},
		},
		NegativeCache: repository.NegativeCache{
			Enabled: true,
			TTL:     1440,
		},
		Proxy: repository.Proxy{
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
			RemoteURL:      "https://archive.ubuntu.com/ubuntu/",
		},
		Storage: repository.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
	}
}

func TestAptProxyRepository(t *testing.T) {
	service := getTestService()
	routingRuleService := nexus3.NewRoutingRuleService(getTestClient())
	routingRule := schema.RoutingRule{
		Name:        strconv.Itoa(rand.Intn(1024)),
		Description: "test",
		Mode:        schema.RoutingRuleModeAllow,
		Matchers: []string{
			"/",
		},
	}
	err := routingRuleService.Create(&routingRule)
	defer routingRuleService.Delete(routingRule.Name)
	assert.Nil(t, err)
	repo := getTestAptProxyRepository("test-apt-repo-hosted-" + strconv.Itoa(rand.Intn(1024)))
	repo.RoutingRule = &routingRule.Name

	err = service.Proxy.Create(repo)
	assert.Nil(t, err)
	generatedRepo, err := service.Proxy.Get(repo.Name)
	assert.Nil(t, err)
	assert.Equal(t, repo.Online, generatedRepo.Online)
	assert.Equal(t, repo.Apt, generatedRepo.Apt)
	assert.Equal(t, repo.HTTPClient.Blocked, generatedRepo.HTTPClient.Blocked)
	assert.Equal(t, repo.HTTPClient.AutoBlock, generatedRepo.HTTPClient.AutoBlock)
	assert.Equal(t, repo.HTTPClient.Connection.Timeout, generatedRepo.HTTPClient.Connection.Timeout)
	assert.Equal(t, repo.HTTPClient.Connection.UseTrustStore, generatedRepo.HTTPClient.Connection.UseTrustStore)
	assert.Equal(t, repo.NegativeCache, generatedRepo.NegativeCache)
	assert.Equal(t, repo.Proxy, generatedRepo.Proxy)
	assert.Equal(t, repo.RoutingRule, generatedRepo.RoutingRuleName)
	assert.Equal(t, repo.Storage, generatedRepo.Storage)

	updatedRepo := repo
	updatedRepo.Online = false
	updatedRepo.Apt.Flat = false

	err = service.Proxy.Update(repo.Name, updatedRepo)
	assert.Nil(t, err)
	generatedRepo, err = service.Proxy.Get(updatedRepo.Name)
	assert.Nil(t, err)
	assert.Equal(t, updatedRepo.Online, generatedRepo.Online)
	assert.Equal(t, updatedRepo.Apt, generatedRepo.Apt)

	service.Proxy.Delete(repo.Name)
	assert.Nil(t, err)
}
