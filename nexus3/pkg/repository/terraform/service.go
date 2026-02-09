package terraform

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)


type RepositoryTerraformProxyService =
	common.RepositoryService[repository.TerraformProxyRepository]

type RepositoryTerraformHostedService =
	common.RepositoryService[repository.TerraformHostedRepository]

// NewRepositoryTerraformProxyService creates a new Terraform proxy repository service.
func NewRepositoryTerraformProxyService(c *client.Client) *RepositoryTerraformProxyService {
	endpoint := common.RepositoryAPIEndpoint + "/terraform/proxy"
	return common.NewRepositoryService[repository.TerraformProxyRepository](endpoint, c)
}

func NewRepositoryTerraformHostedService(c *client.Client) *RepositoryTerraformHostedService {
	endpoint := common.RepositoryAPIEndpoint + "/terraform/hosted"
	return common.NewRepositoryService[repository.TerraformHostedRepository](endpoint, c)
}

// RepositoryTerraformService groups all Terraform repository services.
type RepositoryTerraformService struct {
	Proxy   *RepositoryTerraformProxyService
	Hosted  *RepositoryTerraformHostedService
}

// NewRepositoryTerraformService creates a new Terraform repository service group.
func NewRepositoryTerraformService(c *client.Client) *RepositoryTerraformService {
	return &RepositoryTerraformService{
		Proxy:  NewRepositoryTerraformProxyService(c),
		Hosted: NewRepositoryTerraformHostedService(c),
	}
}