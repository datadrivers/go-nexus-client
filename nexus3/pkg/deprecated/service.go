package deprecated

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
)

const (
	securityAPIEndpoint = client.BasePath + "v1/security"
)

type DeprecatedService struct {
	client *client.Client

	// API Services
	Privilege *SecurityPrivilegeService
}

func NewDeprecatedService(c *client.Client) *DeprecatedService {
	return &DeprecatedService{
		client: c,

		Privilege: NewSecurityPrivilegeService(c),
	}
}
