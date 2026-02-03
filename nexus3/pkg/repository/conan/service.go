package conan

import (
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

const (
	conanAPIEndpoint = common.RepositoryAPIEndpoint + "/conan"
)

type (
	RepositoryConanProxyService = common.RepositoryService[repository.ConanProxyRepository]
)

type RepositoryConanService struct {
	client *client.Client

	Proxy *RepositoryConanProxyService
}

func NewRepositoryConanService(c *client.Client) *RepositoryConanService {
	return &RepositoryConanService{
		client: c,

		Proxy: NewRepositoryConanProxyService(c),
	}
}

func NewRepositoryConanProxyService(c *client.Client) *RepositoryConanProxyService {
	return common.NewRepositoryService[repository.ConanProxyRepository](conanAPIEndpoint+"/proxy", c)
}
