package pypi

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	pypiAPIEndpoint = common.RepositoryAPIEndpoint + "/pypi"
)

type (
	RepositoryPypiGroupService  = common.RepositoryService[repository.PypiGroupRepository]
	RepositoryPypiHostedService = common.RepositoryService[repository.PypiHostedRepository]
	RepositoryPypiProxyService  = common.RepositoryService[repository.PypiProxyRepository]
)

type RepositoryPypiService struct {
	client *client.Client

	Group  *RepositoryPypiGroupService
	Hosted *RepositoryPypiHostedService
	Proxy  *RepositoryPypiProxyService
}

func NewRepositoryPypiService(c *client.Client) *RepositoryPypiService {
	return &RepositoryPypiService{
		client: c,

		Group:  NewRepositoryPypiGroupService(c),
		Hosted: NewRepositoryPypiHostedService(c),
		Proxy:  NewRepositoryPypiProxyService(c),
	}
}

func NewRepositoryPypiGroupService(c *client.Client) *RepositoryPypiGroupService {
	return common.NewRepositoryService[repository.PypiGroupRepository](pypiAPIEndpoint+"/group", c)
}

func NewRepositoryPypiHostedService(c *client.Client) *RepositoryPypiHostedService {
	return common.NewRepositoryService[repository.PypiHostedRepository](pypiAPIEndpoint+"/hosted", c)
}

func NewRepositoryPypiProxyService(c *client.Client) *RepositoryPypiProxyService {
	return common.NewRepositoryService[repository.PypiProxyRepository](pypiAPIEndpoint+"/proxy", c)
}
