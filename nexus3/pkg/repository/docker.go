package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
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
