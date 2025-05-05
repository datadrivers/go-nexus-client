package golang

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	goAPIEndpoint = common.RepositoryAPIEndpoint + "/go"
)

type (
	RepositoryGoGroupService = common.RepositoryService[repository.GoGroupRepository]
	RepositoryGoProxyService = common.RepositoryService[repository.GoProxyRepository]
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

func NewRepositoryGoGroupService(c *client.Client) *RepositoryGoGroupService {
	return common.NewRepositoryService[repository.GoGroupRepository](goAPIEndpoint+"/group", c)
}

func NewRepositoryGoProxyService(c *client.Client) *RepositoryGoProxyService {
	return common.NewRepositoryService[repository.GoProxyRepository](goAPIEndpoint+"/proxy", c)
}
