package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
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
