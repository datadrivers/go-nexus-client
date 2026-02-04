package terraform

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

// RepositoryTerraformProxyService provides access to
// Terraform proxy repository REST endpoints:
//
//   GET    /v1/repositories/terraform/proxy/{repositoryName}
//   POST   /v1/repositories/terraform/proxy
//   PUT    /v1/repositories/terraform/proxy/{repositoryName}
type RepositoryTerraformProxyService =
	common.RepositoryService[repository.TerraformProxyRepository]

// NewRepositoryTerraformProxyService creates a new Terraform proxy repository service.
func NewRepositoryTerraformProxyService(c *client.Client) *RepositoryTerraformProxyService {
	endpoint := common.RepositoryAPIEndpoint + "/terraform/proxy"
	return common.NewRepositoryService[repository.TerraformProxyRepository](endpoint, c)
}

// RepositoryTerraformService groups all Terraform repository services.
type RepositoryTerraformService struct {
	Proxy *RepositoryTerraformProxyService
}

// NewRepositoryTerraformService creates a new Terraform repository service group.
func NewRepositoryTerraformService(c *client.Client) *RepositoryTerraformService {
	return &RepositoryTerraformService{
		Proxy: NewRepositoryTerraformProxyService(c),
	}
}
