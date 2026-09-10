package settings

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
)

const (
	settingsAPIEndpoint = client.BasePath + "v1"
)

// SettingsService groups the Nexus system settings related services
type SettingsService struct {
	client *client.Client

	// API Services
	HTTP *SettingsHTTPService
}

// NewSettingsService creates a new instance of SettingsService
func NewSettingsService(c *client.Client) *SettingsService {
	return &SettingsService{
		client: c,

		HTTP: NewSettingsHTTPService(c),
	}
}
