package apt

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const aptAPIEndpoint = common.RepositoryAPIEndpoint + "/apt"

type (
	RepositoryAptHostedService = common.RepositoryService[repository.AptHostedRepository]
	RepositoryAptProxyService  = common.RepositoryService[repository.AptProxyRepository]
)

type RepositoryAptService struct {
	client *client.Client

	Hosted *RepositoryAptHostedService
	Proxy  *RepositoryAptProxyService
}

func NewRepositoryAptService(c *client.Client) *RepositoryAptService {
	return &RepositoryAptService{
		client: c,

		Hosted: NewRepositoryAptHostedService(c),
		Proxy:  NewRepositoryAptProxyService(c),
	}
}

func NewRepositoryAptHostedService(c *client.Client) *RepositoryAptHostedService {
	return common.NewRepositoryService[repository.AptHostedRepository](aptAPIEndpoint+"/hosted", c)
}

func NewRepositoryAptProxyService(c *client.Client) *RepositoryAptProxyService {
	return common.NewRepositoryService[repository.AptProxyRepository](aptAPIEndpoint+"/proxy", c)
}
