package cocoapods

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
)

const (
	cocoapodsAPIEndpoint = common.RepositoryAPIEndpoint + "/cocoapods"
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
