package helm

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	helmAPIEndpoint = common.RepositoryAPIEndpoint + "/helm"
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
