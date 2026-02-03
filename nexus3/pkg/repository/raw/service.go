package raw

import (
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

const (
	rawAPIEndpoint = common.RepositoryAPIEndpoint + "/raw"
)

type (
	RepositoryRawGroupService  = common.RepositoryService[repository.RawGroupRepository]
	RepositoryRawHostedService = common.RepositoryService[repository.RawHostedRepository]
	RepositoryRawProxyService  = common.RepositoryService[repository.RawProxyRepository]
)

type RepositoryRawService struct {
	client *client.Client

	Group  *RepositoryRawGroupService
	Hosted *RepositoryRawHostedService
	Proxy  *RepositoryRawProxyService
}

func NewRepositoryRawService(c *client.Client) *RepositoryRawService {
	return &RepositoryRawService{
		client: c,

		Group:  NewRepositoryRawGroupService(c),
		Hosted: NewRepositoryRawHostedService(c),
		Proxy:  NewRepositoryRawProxyService(c),
	}
}

func NewRepositoryRawGroupService(c *client.Client) *RepositoryRawGroupService {
	return common.NewRepositoryService[repository.RawGroupRepository](rawAPIEndpoint+"/group", c)
}

func NewRepositoryRawHostedService(c *client.Client) *RepositoryRawHostedService {
	return common.NewRepositoryService[repository.RawHostedRepository](rawAPIEndpoint+"/hosted", c)
}

func NewRepositoryRawProxyService(c *client.Client) *RepositoryRawProxyService {
	return common.NewRepositoryService[repository.RawProxyRepository](rawAPIEndpoint+"/proxy", c)
}
