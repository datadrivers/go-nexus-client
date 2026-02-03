package cargo

import (
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

const cargoAPIEndpoint = common.RepositoryAPIEndpoint + "/cargo"

type (
	RepositoryCargoGroupService  = common.RepositoryService[repository.CargoGroupRepository]
	RepositoryCargoHostedService = common.RepositoryService[repository.CargoHostedRepository]
	RepositoryCargoProxyService  = common.RepositoryService[repository.CargoProxyRepository]
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

func NewRepositoryCargoGroupService(c *client.Client) *RepositoryCargoGroupService {
	return common.NewRepositoryService[repository.CargoGroupRepository](cargoAPIEndpoint+"/group", c)
}

func NewRepositoryCargoHostedService(c *client.Client) *RepositoryCargoHostedService {
	return common.NewRepositoryService[repository.CargoHostedRepository](cargoAPIEndpoint+"/hosted", c)
}

func NewRepositoryCargoProxyService(c *client.Client) *RepositoryCargoProxyService {
	return common.NewRepositoryService[repository.CargoProxyRepository](cargoAPIEndpoint+"/proxy", c)
}
