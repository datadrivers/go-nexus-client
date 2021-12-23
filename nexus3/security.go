package nexus3

const (
	securityAPIEndpoint = basePath + "v1/security"
)

type SecurityService struct {
	client *client

	// API Services
	Anonymous       *SecurityAnonymousService
	ContentSelector *SecurityContentSelectorService
	LDAP            *SecurityLdapService
	Privilege       *SecurityPrivilegeService
	SSL             *SecuritySSLService
	User            *SecurityUserService
	UserTokens      *SecurityUserTokensService
}

func NewSecurityService(c *client) *SecurityService {
	return &SecurityService{
		client: c,

		Anonymous:       NewSecurityAnonymousService(c),
		ContentSelector: NewSecurityContentSelectorService(c),
		LDAP:            NewSecurityLdapService(c),
		Privilege:       NewSecurityPrivilegeService(c),
		SSL:             NewSecuritySSLService(c),
		User:            NewSecurityUserService(c),
		UserTokens:      NewSecurityUserTokensService(c),
	}
}
