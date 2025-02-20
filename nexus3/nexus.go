package nexus3

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/blobstore"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/readonly"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/security"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/task"
)

const (
	// ContentTypeApplicationJSON ...
	ContentTypeApplicationJSON = "application/json"
	// ContentTypeTextPlain ...
	ContentTypeTextPlain = "text/plain"
	basePath             = "service/rest/"
)

type NexusClient struct {
	client *client.Client

	// API Services
	BlobStore   *blobstore.BlobStoreService
	Repository  *repository.RepositoryService
	RoutingRule *RoutingRuleService
	Security    *security.SecurityService
	Script      *ScriptService
	ReadOnly    *readonly.ReadOnlyService
	MailConfig  *MailConfigService
	Task        *task.TaskService
}

// NewClient returns an instance of client that implements the Client interface
func NewClient(config client.Config) *NexusClient {
	client := client.NewClient(config)
	return &NexusClient{
		client:      client,
		BlobStore:   blobstore.NewBlobStoreService(client),
		Repository:  repository.NewRepositoryService(client),
		RoutingRule: NewRoutingRuleService(client),
		Security:    security.NewSecurityService(client),
		Script:      NewScriptService(client),
		ReadOnly:    readonly.NewReadOnlyService(client),
		MailConfig:  NewMailConfigService(client),
		Task:        task.NewTaskService(client),
	}
}
