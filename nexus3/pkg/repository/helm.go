package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
)

type RepositoryHelmService struct {
	client *client.Client

	Proxy *RepositoryHelmProxyService
}

func NewRepositoryHelmService(c *client.Client) *RepositoryHelmService {
	return &RepositoryHelmService{
		client: c,

		Proxy: NewRepositoryHelmProxyService(c),
	}
}
