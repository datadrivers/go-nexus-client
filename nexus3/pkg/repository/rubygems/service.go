package rubygems

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	rubyGemsAPIEndpoint = common.RepositoryAPIEndpoint + "/rubygems"
)

type (
	RepositoryRubyGemsGroupService  = common.RepositoryService[repository.RubyGemsGroupRepository]
	RepositoryRubyGemsHostedService = common.RepositoryService[repository.RubyGemsHostedRepository]
	RepositoryRubyGemsProxyService  = common.RepositoryService[repository.RubyGemsProxyRepository]
)

type RepositoryRubyGemsService struct {
	client *client.Client

	Group  *RepositoryRubyGemsGroupService
	Hosted *RepositoryRubyGemsHostedService
	Proxy  *RepositoryRubyGemsProxyService
}

func NewRepositoryRubyGemsService(c *client.Client) *RepositoryRubyGemsService {
	return &RepositoryRubyGemsService{
		client: c,

		Group:  NewRepositoryRubyGemsGroupService(c),
		Hosted: NewRepositoryRubyGemsHostedService(c),
		Proxy:  NewRepositoryRubyGemsProxyService(c),
	}
}

func NewRepositoryRubyGemsGroupService(c *client.Client) *RepositoryRubyGemsGroupService {
	return common.NewRepositoryService[repository.RubyGemsGroupRepository](rubyGemsAPIEndpoint+"/group", c)
}

func NewRepositoryRubyGemsHostedService(c *client.Client) *RepositoryRubyGemsHostedService {
	return common.NewRepositoryService[repository.RubyGemsHostedRepository](rubyGemsAPIEndpoint+"/hosted", c)
}

func NewRepositoryRubyGemsProxyService(c *client.Client) *RepositoryRubyGemsProxyService {
	return common.NewRepositoryService[repository.RubyGemsProxyRepository](rubyGemsAPIEndpoint+"/proxy", c)
}
