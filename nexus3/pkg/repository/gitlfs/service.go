package gitlfs

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	gitLfsAPIEndpoint = common.RepositoryAPIEndpoint + "/gitlfs"
)

type (
	RepositoryGitLfsHostedService = common.RepositoryService[repository.GitLfsHostedRepository]
)

type RepositoryGitLfsService struct {
	client *client.Client

	Hosted *RepositoryGitLfsHostedService
}

func NewRepositoryGitLfsService(c *client.Client) *RepositoryGitLfsService {
	return &RepositoryGitLfsService{
		client: c,

		Hosted: NewRepositoryGitLfsHostedService(c),
	}
}

func NewRepositoryGitLfsHostedService(c *client.Client) *RepositoryGitLfsHostedService {
	return common.NewRepositoryService[repository.GitLfsHostedRepository](gitLfsAPIEndpoint+"/hosted", c)
}
