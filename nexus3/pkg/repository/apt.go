package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
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
