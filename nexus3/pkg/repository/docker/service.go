package docker

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	dockerAPIEndpoint = common.RepositoryAPIEndpoint + "/docker"
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
