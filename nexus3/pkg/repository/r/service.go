package r

import (
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

const (
	rAPIEndpoint = common.RepositoryAPIEndpoint + "/r"
)

type (
	RepositoryRGroupService  = common.RepositoryService[repository.RGroupRepository]
	RepositoryRHostedService = common.RepositoryService[repository.RHostedRepository]
	RepositoryRProxyService  = common.RepositoryService[repository.RProxyRepository]
)

type RepositoryRService struct {
	client *client.Client

	Group  *RepositoryRGroupService
	Hosted *RepositoryRHostedService
	Proxy  *RepositoryRProxyService
}

func NewRepositoryRService(c *client.Client) *RepositoryRService {
	return &RepositoryRService{
		client: c,

		Group:  NewRepositoryRGroupService(c),
		Hosted: NewRepositoryRHostedService(c),
		Proxy:  NewRepositoryRProxyService(c),
	}
}

func NewRepositoryRGroupService(c *client.Client) *RepositoryRGroupService {
	return common.NewRepositoryService[repository.RGroupRepository](rAPIEndpoint+"/group", c)
}

func NewRepositoryRHostedService(c *client.Client) *RepositoryRHostedService {
	return common.NewRepositoryService[repository.RHostedRepository](rAPIEndpoint+"/hosted", c)
}

func NewRepositoryRProxyService(c *client.Client) *RepositoryRProxyService {
	return common.NewRepositoryService[repository.RProxyRepository](rAPIEndpoint+"/proxy", c)
}
