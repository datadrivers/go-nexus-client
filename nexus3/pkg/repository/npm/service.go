package npm

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	npmAPIEndpoint = common.RepositoryAPIEndpoint + "/npm"
)

type (
	RepositoryNpmGroupService  = common.RepositoryService[repository.NpmGroupRepository]
	RepositoryNpmHostedService = common.RepositoryService[repository.NpmHostedRepository]
	RepositoryNpmProxyService  = common.RepositoryService[repository.NpmProxyRepository]
)

type RepositoryNpmService struct {
	client *client.Client

	Group  *RepositoryNpmGroupService
	Hosted *RepositoryNpmHostedService
	Proxy  *RepositoryNpmProxyService
}

func NewRepositoryNpmService(c *client.Client) *RepositoryNpmService {
	return &RepositoryNpmService{
		client: c,

		Group:  NewRepositoryNpmGroupService(c),
		Hosted: NewRepositoryNpmHostedService(c),
		Proxy:  NewRepositoryNpmProxyService(c),
	}
}

func NewRepositoryNpmGroupService(c *client.Client) *RepositoryNpmGroupService {
	return common.NewRepositoryService[repository.NpmGroupRepository](npmAPIEndpoint+"/group", c)
}

func NewRepositoryNpmHostedService(c *client.Client) *RepositoryNpmHostedService {
	return common.NewRepositoryService[repository.NpmHostedRepository](npmAPIEndpoint+"/hosted", c)
}

func NewRepositoryNpmProxyService(c *client.Client) *RepositoryNpmProxyService {
	return common.NewRepositoryService[repository.NpmProxyRepository](npmAPIEndpoint+"/proxy", c)
}
