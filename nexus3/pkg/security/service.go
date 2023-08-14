package security

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/security/privilege"
)

const (
	securityAPIEndpoint = client.BasePath + "v1/security"
)

type SecurityService struct {
	client *client.Client

	// API Services
	Anonymous       *SecurityAnonymousService
	ContentSelector *SecurityContentSelectorService
	LDAP            *SecurityLdapService
	Privilege       *privilege.SecurityPrivilegeService
	Realm           *SecurityRealmService
	Role            *SecurityRoleService
	SAML            *SecuritySamlService
	SSL             *SecuritySSLService
	User            *SecurityUserService
	UserTokens      *SecurityUserTokensService
}

func NewSecurityService(c *client.Client) *SecurityService {
	return &SecurityService{
		client: c,

		Anonymous:       NewSecurityAnonymousService(c),
		ContentSelector: NewSecurityContentSelectorService(c),
		LDAP:            NewSecurityLdapService(c),
		Privilege:       privilege.NewSecurityPrivilegeService(c),
		Realm:           NewSecurityRealmService(c),
		Role:            NewSecurityRoleService(c),
		SAML:            NewSecuritySamlService(c),
		SSL:             NewSecuritySSLService(c),
		User:            NewSecurityUserService(c),
		UserTokens:      NewSecurityUserTokensService(c),
	}
}
