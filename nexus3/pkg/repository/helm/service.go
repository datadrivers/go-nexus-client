package helm

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	helmAPIEndpoint = common.RepositoryAPIEndpoint + "/helm"
)

type (
	RepositoryHelmHostedService = common.RepositoryService[repository.HelmHostedRepository]
	RepositoryHelmProxyService  = common.RepositoryService[repository.HelmProxyRepository]
)

type RepositoryHelmService struct {
	client *client.Client

	Hosted *RepositoryHelmHostedService
	Proxy  *RepositoryHelmProxyService
}

func NewRepositoryHelmService(c *client.Client) *RepositoryHelmService {
	return &RepositoryHelmService{
		client: c,

		Hosted: NewRepositoryHelmHostedService(c),
		Proxy:  NewRepositoryHelmProxyService(c),
	}
}

func NewRepositoryHelmHostedService(c *client.Client) *RepositoryHelmHostedService {
	return common.NewRepositoryService[repository.HelmHostedRepository](helmAPIEndpoint+"/hosted", c)
}

func NewRepositoryHelmProxyService(c *client.Client) *RepositoryHelmProxyService {
	return common.NewRepositoryService[repository.HelmProxyRepository](helmAPIEndpoint+"/proxy", c)
}
