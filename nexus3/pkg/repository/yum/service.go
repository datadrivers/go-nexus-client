package yum

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	yumAPIEndpoint = common.RepositoryAPIEndpoint + "/yum"
)

type RepositoryYumService struct {
	client *client.Client

	Group  *RepositoryYumGroupService
	Hosted *RepositoryYumHostedService
	Proxy  *RepositoryYumProxyService
}

func NewRepositoryYumService(c *client.Client) *RepositoryYumService {
	return &RepositoryYumService{
		client: c,

		Group:  NewRepositoryYumGroupService(c),
		Hosted: NewRepositoryYumHostedService(c),
		Proxy:  NewRepositoryYumProxyService(c),
	}
}
