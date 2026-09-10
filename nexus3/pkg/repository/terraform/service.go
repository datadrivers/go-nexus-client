package terraform

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	terraformAPIEndpoint = common.RepositoryAPIEndpoint + "/terraform"
)

type (
	RepositoryTerraformGroupService  = common.RepositoryService[repository.TerraformGroupRepository]
	RepositoryTerraformHostedService = common.RepositoryService[repository.TerraformHostedRepository]
	RepositoryTerraformProxyService  = common.RepositoryService[repository.TerraformProxyRepository]
)

type RepositoryTerraformService struct {
	client *client.Client

	Group  *RepositoryTerraformGroupService
	Hosted *RepositoryTerraformHostedService
	Proxy  *RepositoryTerraformProxyService
}

func NewRepositoryTerraformService(c *client.Client) *RepositoryTerraformService {
	return &RepositoryTerraformService{
		client: c,

		Group:  NewRepositoryTerraformGroupService(c),
		Hosted: NewRepositoryTerraformHostedService(c),
		Proxy:  NewRepositoryTerraformProxyService(c),
	}
}

func NewRepositoryTerraformGroupService(c *client.Client) *RepositoryTerraformGroupService {
	return common.NewRepositoryService[repository.TerraformGroupRepository](terraformAPIEndpoint+"/group", c)
}

func NewRepositoryTerraformHostedService(c *client.Client) *RepositoryTerraformHostedService {
	return common.NewRepositoryService[repository.TerraformHostedRepository](terraformAPIEndpoint+"/hosted", c)
}

func NewRepositoryTerraformProxyService(c *client.Client) *RepositoryTerraformProxyService {
	return common.NewRepositoryService[repository.TerraformProxyRepository](terraformAPIEndpoint+"/proxy", c)
}
