package alpine

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const alpineAPIEndpoint = common.RepositoryAPIEndpoint + "/alpine"

type RepositoryAlpineProxyService = common.RepositoryService[repository.AlpineProxyRepository]

type RepositoryAlpineService struct {
	client *client.Client

	Proxy *RepositoryAlpineProxyService
}

func NewRepositoryAlpineService(c *client.Client) *RepositoryAlpineService {
	return &RepositoryAlpineService{
		client: c,

		Proxy: NewRepositoryAlpineProxyService(c),
	}
}

func NewRepositoryAlpineProxyService(c *client.Client) *RepositoryAlpineProxyService {
	return common.NewRepositoryService[repository.AlpineProxyRepository](alpineAPIEndpoint+"/proxy", c)
}
