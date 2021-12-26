package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
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
