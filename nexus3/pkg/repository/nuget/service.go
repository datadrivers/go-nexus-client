package nuget

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	nugetAPIEndpoint = common.RepositoryAPIEndpoint + "/nuget"
)

type (
	RepositoryNugetGroupService  = common.RepositoryService[repository.NugetGroupRepository]
	RepositoryNugetHostedService = common.RepositoryService[repository.NugetHostedRepository]
	RepositoryNugetProxyService  = common.RepositoryService[repository.NugetProxyRepository]
)

type RepositoryNugetService struct {
	client *client.Client

	Group  *RepositoryNugetGroupService
	Hosted *RepositoryNugetHostedService
	Proxy  *RepositoryNugetProxyService
}

func NewRepositoryNugetService(c *client.Client) *RepositoryNugetService {
	return &RepositoryNugetService{
		client: c,

		Group:  NewRepositoryNugetGroupService(c),
		Hosted: NewRepositoryNugetHostedService(c),
		Proxy:  NewRepositoryNugetProxyService(c),
	}
}

func NewRepositoryNugetGroupService(c *client.Client) *RepositoryNugetGroupService {
	return common.NewRepositoryService[repository.NugetGroupRepository](nugetAPIEndpoint+"/group", c)
}

func NewRepositoryNugetHostedService(c *client.Client) *RepositoryNugetHostedService {
	return common.NewRepositoryService[repository.NugetHostedRepository](nugetAPIEndpoint+"/hosted", c)
}

func NewRepositoryNugetProxyService(c *client.Client) *RepositoryNugetProxyService {
	return common.NewRepositoryService[repository.NugetProxyRepository](nugetAPIEndpoint+"/proxy", c)
}
