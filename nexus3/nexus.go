package nexus3

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/blobstore"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/security"
)

const (
	// ContentTypeApplicationJSON ...
	ContentTypeApplicationJSON = "application/json"
	// ContentTypeTextPlain ...
	ContentTypeTextPlain = "text/plain"
	basePath             = "service/rest/"
)

type nexusClient struct {
	client *client.Client

	// API Services
	BlobStore   *blobstore.BlobStoreService
	Repository  *repository.RepositoryService
	RoutingRule *RoutingRuleService
	Security    *security.SecurityService
	Script      *ScriptService
}

// NewClient returns an instance of client that implements the Client interface
func NewClient(config client.Config) *nexusClient {
	client := client.NewClient(config)
	return &nexusClient{
		client:      client,
		BlobStore:   blobstore.NewBlobStoreService(client),
		Repository:  repository.NewRepositoryService(client),
		RoutingRule: NewRoutingRuleService(client),
		Security:    security.NewSecurityService(client),
		Script:      NewScriptService(client),
	}
}
