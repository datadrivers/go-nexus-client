package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
)

type RepositoryCocoapodsService struct {
	client *client.Client

	Proxy *RepositoryCocoapodsProxyService
}

func NewRepositoryCocoapodsService(c *client.Client) *RepositoryCocoapodsService {
	return &RepositoryCocoapodsService{
		client: c,

		Proxy: NewRepositoryCocoapodsProxyService(c),
	}
}
