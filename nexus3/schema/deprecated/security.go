package deprecated

import "github.com/datadrivers/go-nexus-client/nexus3/pkg/client"

type SecurityPrivilegeService client.Service

type SecurityService struct {
	client    *client.Client
	Privilege *SecurityPrivilegeService
}
