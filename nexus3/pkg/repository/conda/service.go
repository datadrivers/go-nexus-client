package conda

import (
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

const (
	condaAPIEndpoint = common.RepositoryAPIEndpoint + "/conda"
)

type (
	RepositoryCondaProxyService = common.RepositoryService[repository.CondaProxyRepository]
)

type RepositoryCondaService struct {
	client *client.Client

	Proxy *RepositoryCondaProxyService
}

func NewRepositoryCondaService(c *client.Client) *RepositoryCondaService {
	return &RepositoryCondaService{
		client: c,

		Proxy: NewRepositoryCondaProxyService(c),
	}
}

func NewRepositoryCondaProxyService(c *client.Client) *RepositoryCondaProxyService {
	return common.NewRepositoryService[repository.CondaProxyRepository](condaAPIEndpoint+"/proxy", c)
}
