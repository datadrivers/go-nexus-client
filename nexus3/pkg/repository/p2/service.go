package p2

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	p2APIEndpoint = common.RepositoryAPIEndpoint + "/p2"
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
