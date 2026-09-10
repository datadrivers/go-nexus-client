package alpine

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	alpineAPIEndpoint = common.RepositoryAPIEndpoint + "/alpine"
)

type (
	RepositoryAlpineGroupService  = common.RepositoryService[repository.AlpineGroupRepository]
	RepositoryAlpineHostedService = common.RepositoryService[repository.AlpineHostedRepository]
	RepositoryAlpineProxyService  = common.RepositoryService[repository.AlpineProxyRepository]
)

type RepositoryAlpineService struct {
	client *client.Client

	Group  *RepositoryAlpineGroupService
	Hosted *RepositoryAlpineHostedService
	Proxy  *RepositoryAlpineProxyService
}

func NewRepositoryAlpineService(c *client.Client) *RepositoryAlpineService {
	return &RepositoryAlpineService{
		client: c,

		Group:  NewRepositoryAlpineGroupService(c),
		Hosted: NewRepositoryAlpineHostedService(c),
		Proxy:  NewRepositoryAlpineProxyService(c),
	}
}

func NewRepositoryAlpineGroupService(c *client.Client) *RepositoryAlpineGroupService {
	return common.NewRepositoryService[repository.AlpineGroupRepository](alpineAPIEndpoint+"/group", c)
}

func NewRepositoryAlpineHostedService(c *client.Client) *RepositoryAlpineHostedService {
	return common.NewRepositoryService[repository.AlpineHostedRepository](alpineAPIEndpoint+"/hosted", c)
}

func NewRepositoryAlpineProxyService(c *client.Client) *RepositoryAlpineProxyService {
	return common.NewRepositoryService[repository.AlpineProxyRepository](alpineAPIEndpoint+"/proxy", c)
}
