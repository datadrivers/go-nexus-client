package conan

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	conanAPIEndpoint = common.RepositoryAPIEndpoint + "/conan"
)

type (
	RepositoryConanProxyService  = common.RepositoryService[repository.ConanProxyRepository]
	RepositoryConanHostedService = common.RepositoryService[repository.ConanHostedRepository]
)

type RepositoryConanService struct {
	client *client.Client

	Proxy  *RepositoryConanProxyService
	Hosted *RepositoryConanHostedService
}

func NewRepositoryConanService(c *client.Client) *RepositoryConanService {
	return &RepositoryConanService{
		client: c,
		Proxy:  NewRepositoryConanProxyService(c),
		Hosted: NewRepositoryConanHostedService(c),
	}
}

func NewRepositoryConanProxyService(c *client.Client) *RepositoryConanProxyService {
	return common.NewRepositoryService[repository.ConanProxyRepository](conanAPIEndpoint+"/proxy", c)
}

func NewRepositoryConanHostedService(c *client.Client) *RepositoryConanHostedService {
	return common.NewRepositoryService[repository.ConanHostedRepository](conanAPIEndpoint+"/hosted", c)
}
