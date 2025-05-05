package yum

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	yumAPIEndpoint = common.RepositoryAPIEndpoint + "/yum"
)

type (
	RepositoryYumGroupService  = common.RepositoryService[repository.YumGroupRepository]
	RepositoryYumHostedService = common.RepositoryService[repository.YumHostedRepository]
	RepositoryYumProxyService  = common.RepositoryService[repository.YumProxyRepository]
)

type RepositoryYumService struct {
	client *client.Client

	Group  *RepositoryYumGroupService
	Hosted *RepositoryYumHostedService
	Proxy  *RepositoryYumProxyService
}

func NewRepositoryYumService(c *client.Client) *RepositoryYumService {
	return &RepositoryYumService{
		client: c,

		Group:  NewRepositoryYumGroupService(c),
		Hosted: NewRepositoryYumHostedService(c),
		Proxy:  NewRepositoryYumProxyService(c),
	}
}
func NewRepositoryYumGroupService(c *client.Client) *RepositoryYumGroupService {
	return common.NewRepositoryService[repository.YumGroupRepository](yumAPIEndpoint+"/group", c)
}

func NewRepositoryYumHostedService(c *client.Client) *RepositoryYumHostedService {
	return common.NewRepositoryService[repository.YumHostedRepository](yumAPIEndpoint+"/hosted", c)
}

func NewRepositoryYumProxyService(c *client.Client) *RepositoryYumProxyService {
	return common.NewRepositoryService[repository.YumProxyRepository](yumAPIEndpoint+"/proxy", c)
}
