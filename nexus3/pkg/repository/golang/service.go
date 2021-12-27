package golang

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	goAPIEndpoint = common.RepositoryAPIEndpoint + "/go"
)

type RepositoryGoService struct {
	client *client.Client

	Group *RepositoryGoGroupService
	Proxy *RepositoryGoProxyService
}

func NewRepositoryGoService(c *client.Client) *RepositoryGoService {
	return &RepositoryGoService{
		client: c,

		Group: NewRepositoryGoGroupService(c),
		Proxy: NewRepositoryGoProxyService(c),
	}
}
