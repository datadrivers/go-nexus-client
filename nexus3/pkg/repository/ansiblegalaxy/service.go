package ansiblegalaxy

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	ansibleGalaxyAPIEndpoint = common.RepositoryAPIEndpoint + "/ansiblegalaxy"
)

type (
	RepositoryAnsibleGalaxyGroupService  = common.RepositoryService[repository.AnsibleGalaxyGroupRepository]
	RepositoryAnsibleGalaxyHostedService = common.RepositoryService[repository.AnsibleGalaxyHostedRepository]
	RepositoryAnsibleGalaxyProxyService  = common.RepositoryService[repository.AnsibleGalaxyProxyRepository]
)

type RepositoryAnsibleGalaxyService struct {
	client *client.Client

	Group  *RepositoryAnsibleGalaxyGroupService
	Hosted *RepositoryAnsibleGalaxyHostedService
	Proxy  *RepositoryAnsibleGalaxyProxyService
}

func NewRepositoryAnsibleGalaxyService(c *client.Client) *RepositoryAnsibleGalaxyService {
	return &RepositoryAnsibleGalaxyService{
		client: c,

		Group:  NewRepositoryAnsibleGalaxyGroupService(c),
		Hosted: NewRepositoryAnsibleGalaxyHostedService(c),
		Proxy:  NewRepositoryAnsibleGalaxyProxyService(c),
	}
}

func NewRepositoryAnsibleGalaxyGroupService(c *client.Client) *RepositoryAnsibleGalaxyGroupService {
	return common.NewRepositoryService[repository.AnsibleGalaxyGroupRepository](ansibleGalaxyAPIEndpoint+"/group", c)
}

func NewRepositoryAnsibleGalaxyHostedService(c *client.Client) *RepositoryAnsibleGalaxyHostedService {
	return common.NewRepositoryService[repository.AnsibleGalaxyHostedRepository](ansibleGalaxyAPIEndpoint+"/hosted", c)
}

func NewRepositoryAnsibleGalaxyProxyService(c *client.Client) *RepositoryAnsibleGalaxyProxyService {
	return common.NewRepositoryService[repository.AnsibleGalaxyProxyRepository](ansibleGalaxyAPIEndpoint+"/proxy", c)
}
