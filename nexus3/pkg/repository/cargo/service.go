package cargo

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	cargoAPIEndpoint = common.RepositoryAPIEndpoint + "/cargo"
)

type RepositoryCargoService struct {
	client *client.Client

	Group  *RepositoryCargoGroupService
	Proxy  *RepositoryCargoProxyService
	Hosted *RepositoryCargoHostedService
}

func NewRepositoryCargoService(c *client.Client) *RepositoryCargoService {
	return &RepositoryCargoService{
		client: c,

		Group:  NewRepositoryCargoGroupService(c),
		Proxy:  NewRepositoryCargoProxyService(c),
		Hosted: NewRepositoryCargoHostedService(c),
	}
}
