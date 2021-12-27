package nuget

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	nugetAPIEndpoint = common.RepositoryAPIEndpoint + "/nuget"
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
