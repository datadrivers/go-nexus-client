package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
)

type RepositoryConanService struct {
	client *client.Client

	Proxy *RepositoryConanProxyService
}

func NewRepositoryConanService(c *client.Client) *RepositoryConanService {
	return &RepositoryConanService{
		client: c,

		Proxy: NewRepositoryConanProxyService(c),
	}
}
