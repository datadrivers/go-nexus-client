package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
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
