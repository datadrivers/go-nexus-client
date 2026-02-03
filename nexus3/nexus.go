package nexus3

import (
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/cleanup"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/iq"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/readonly"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/repository"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/security"
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
