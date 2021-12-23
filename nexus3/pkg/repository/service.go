package repository

import "github.com/datadrivers/go-nexus-client/nexus3/pkg/client"

const (
	repositoryAPIEndpoint = client.BasePath + "v1/repositories"
)

type RepositoryService struct {
	client *client.Client

	// API Services
}

func NewRepositoryService(c *client.Client) *RepositoryService {
	return &RepositoryService{
		client: c,
	}
}
