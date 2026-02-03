package bower

import (
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

const (
	bowerAPIEndpoint = common.RepositoryAPIEndpoint + "/bower"
)

type (
	RepositoryBowerGroupService  = common.RepositoryService[repository.BowerGroupRepository]
	RepositoryBowerHostedService = common.RepositoryService[repository.BowerHostedRepository]
	RepositoryBowerProxyService  = common.RepositoryService[repository.BowerProxyRepository]
)

type RepositoryBowerService struct {
	client *client.Client

	Group  *RepositoryBowerGroupService
	Hosted *RepositoryBowerHostedService
	Proxy  *RepositoryBowerProxyService
}

func NewRepositoryBowerService(c *client.Client) *RepositoryBowerService {
	return &RepositoryBowerService{
		client: c,

		Group:  NewRepositoryBowerGroupService(c),
		Hosted: NewRepositoryBowerHostedService(c),
		Proxy:  NewRepositoryBowerProxyService(c),
	}
}

func NewRepositoryBowerGroupService(c *client.Client) *RepositoryBowerGroupService {
	return common.NewRepositoryService[repository.BowerGroupRepository](bowerAPIEndpoint+"/group", c)
}

func NewRepositoryBowerHostedService(c *client.Client) *RepositoryBowerHostedService {
	return common.NewRepositoryService[repository.BowerHostedRepository](bowerAPIEndpoint+"/hosted", c)
}

func NewRepositoryBowerProxyService(c *client.Client) *RepositoryBowerProxyService {
	return common.NewRepositoryService[repository.BowerProxyRepository](bowerAPIEndpoint+"/proxy", c)
}
