package pypi

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	pypiAPIEndpoint = common.RepositoryAPIEndpoint + "/pypi"
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
