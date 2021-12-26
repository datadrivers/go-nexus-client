package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
)

type RepositoryGoService struct {
	client *client.Client

	Group *RepositoryGoGroupService
	Proxy *RepositoryGoProxyService
}

func NewRepositoryGoService(c *client.Client) *RepositoryGoService {
	return &RepositoryGoService{
		client: c,

		Group: NewRepositoryGoGroupService(c),
		Proxy: NewRepositoryGoProxyService(c),
	}
}
