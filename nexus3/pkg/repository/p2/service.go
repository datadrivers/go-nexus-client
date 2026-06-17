package p2

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	p2APIEndpoint = common.RepositoryAPIEndpoint + "/p2"
)

type (
	RepositoryP2ProxyService = common.RepositoryService[repository.P2ProxyRepository]
)

type RepositoryP2Service struct {
	client *client.Client

	Proxy *RepositoryP2ProxyService
}

func NewRepositoryP2Service(c *client.Client) *RepositoryP2Service {
	return &RepositoryP2Service{
		client: c,

		Proxy: NewRepositoryP2ProxyService(c),
	}
}

func NewRepositoryP2ProxyService(c *client.Client) *RepositoryP2ProxyService {
	return common.NewRepositoryService[repository.P2ProxyRepository](p2APIEndpoint+"/proxy", c)
}
