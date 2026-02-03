package cocoapods

import (
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

const (
	cocoapodsAPIEndpoint = common.RepositoryAPIEndpoint + "/cocoapods"
)

type (
	RepositoryCocoapodsProxyService = common.RepositoryService[repository.CocoapodsProxyRepository]
)

type RepositoryCocoapodsService struct {
	client *client.Client

	Proxy *RepositoryCocoapodsProxyService
}

func NewRepositoryCocoapodsService(c *client.Client) *RepositoryCocoapodsService {
	return &RepositoryCocoapodsService{
		client: c,

		Proxy: NewRepositoryCocoapodsProxyService(c),
	}
}

func NewRepositoryCocoapodsProxyService(c *client.Client) *RepositoryCocoapodsProxyService {
	return common.NewRepositoryService[repository.CocoapodsProxyRepository](cocoapodsAPIEndpoint+"/proxy", c)
}
