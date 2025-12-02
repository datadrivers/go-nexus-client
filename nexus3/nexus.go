package nexus3

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/blobstore"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/capability"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/cleanup"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/iq"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/readonly"
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

type NexusClient struct {
	client *client.Client

	// API Services
	BlobStore     *blobstore.BlobStoreService
	Capability    *capability.CapabilityService
	CleanupPolicy *cleanup.CleanupPolicyService
	IQServer      *iq.IQServerService
	MailConfig    *MailConfigService
	ReadOnly      *readonly.ReadOnlyService
	Repository    *repository.RepositoryService
	RoutingRule   *RoutingRuleService
	Script        *ScriptService
	Security      *security.SecurityService
}

// NewClient returns an instance of client that implements the Client interface
func NewClient(config client.Config) *NexusClient {
	client := client.NewClient(config)
	return &NexusClient{
		client:        client,
		BlobStore:     blobstore.NewBlobStoreService(client),
		Capability:    capability.NewCapabilityService(client),
		CleanupPolicy: cleanup.NewCleanupPolicyService(client),
		IQServer:      iq.NewIQServerService(client),
		MailConfig:    NewMailConfigService(client),
		ReadOnly:      readonly.NewReadOnlyService(client),
		Repository:    repository.NewRepositoryService(client),
		RoutingRule:   NewRoutingRuleService(client),
		Script:        NewScriptService(client),
		Security:      security.NewSecurityService(client),
	}
}
