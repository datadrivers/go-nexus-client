package npm

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	npmAPIEndpoint = common.RepositoryAPIEndpoint + "/npm"
)

type RepositoryNpmService struct {
	client *client.Client

	Group  *RepositoryNpmGroupService
	Hosted *RepositoryNpmHostedService
	Proxy  *RepositoryNpmProxyService
}

func NewRepositoryNpmService(c *client.Client) *RepositoryNpmService {
	return &RepositoryNpmService{
		client: c,

		Group:  NewRepositoryNpmGroupService(c),
		Hosted: NewRepositoryNpmHostedService(c),
		Proxy:  NewRepositoryNpmProxyService(c),
	}
}
