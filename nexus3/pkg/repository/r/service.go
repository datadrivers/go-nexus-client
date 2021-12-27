package r

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	rAPIEndpoint = common.RepositoryAPIEndpoint + "/r"
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
