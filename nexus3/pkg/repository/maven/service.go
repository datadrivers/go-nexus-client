package maven

import (
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

const (
	mavenAPIEndpoint = common.RepositoryAPIEndpoint + "/maven"
)

type (
	RepositoryMavenGroupService  = common.RepositoryService[repository.MavenGroupRepository]
	RepositoryMavenHostedService = common.RepositoryService[repository.MavenHostedRepository]
	RepositoryMavenProxyService  = common.RepositoryService[repository.MavenProxyRepository]
)

type RepositoryMavenService struct {
	client *client.Client

	Group  *RepositoryMavenGroupService
	Hosted *RepositoryMavenHostedService
	Proxy  *RepositoryMavenProxyService
}

func NewRepositoryMavenService(c *client.Client) *RepositoryMavenService {
	return &RepositoryMavenService{
		client: c,

		Group:  NewRepositoryMavenGroupService(c),
		Hosted: NewRepositoryMavenHostedService(c),
		Proxy:  NewRepositoryMavenProxyService(c),
	}
}

func NewRepositoryMavenGroupService(c *client.Client) *RepositoryMavenGroupService {
	return common.NewRepositoryService[repository.MavenGroupRepository](mavenAPIEndpoint+"/group", c)
}

func NewRepositoryMavenHostedService(c *client.Client) *RepositoryMavenHostedService {
	return common.NewRepositoryService[repository.MavenHostedRepository](mavenAPIEndpoint+"/hosted", c)
}

func NewRepositoryMavenProxyService(c *client.Client) *RepositoryMavenProxyService {
	return common.NewRepositoryService[repository.MavenProxyRepository](mavenAPIEndpoint+"/proxy", c)
}
