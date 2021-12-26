package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
)

type RepositoryP2Service struct {
	client *client.Client

	Proxy *RepositoryP2ProxyService
}

func NewRepositoryP2Service(c *client.Client) *RepositoryP2Service {
	return &RepositoryP2Service{
		client: c,

		Proxy: NewRepositoryP2ProxyService(c),
	}
}
