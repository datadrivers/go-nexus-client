package docker

import (
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

const (
	dockerAPIEndpoint = common.RepositoryAPIEndpoint + "/docker"
)

type (
	RepositoryDockerGroupService  = common.RepositoryService[repository.DockerGroupRepository]
	RepositoryDockerHostedService = common.RepositoryService[repository.DockerHostedRepository]
	RepositoryDockerProxyService  = common.RepositoryService[repository.DockerProxyRepository]
)

type RepositoryDockerService struct {
	client *client.Client

	Group  *RepositoryDockerGroupService
	Hosted *RepositoryDockerHostedService
	Proxy  *RepositoryDockerProxyService
}

func NewRepositoryDockerService(c *client.Client) *RepositoryDockerService {
	return &RepositoryDockerService{
		client: c,

		Group:  NewRepositoryDockerGroupService(c),
		Hosted: NewRepositoryDockerHostedService(c),
		Proxy:  NewRepositoryDockerProxyService(c),
	}
}

func NewRepositoryDockerGroupService(c *client.Client) *RepositoryDockerGroupService {
	return common.NewRepositoryService[repository.DockerGroupRepository](dockerAPIEndpoint+"/group", c)
}

func NewRepositoryDockerHostedService(c *client.Client) *RepositoryDockerHostedService {
	return common.NewRepositoryService[repository.DockerHostedRepository](dockerAPIEndpoint+"/hosted", c)
}

func NewRepositoryDockerProxyService(c *client.Client) *RepositoryDockerProxyService {
	return common.NewRepositoryService[repository.DockerProxyRepository](dockerAPIEndpoint+"/proxy", c)
}
