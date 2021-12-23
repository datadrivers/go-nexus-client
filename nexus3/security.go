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
	SSL             *SecuritySSLService
	User            *SecurityUserService
}

func NewSecurityService(c *client) *SecurityService {
	a := NewSecurityAnonymousService(c)
	cs := NewSecurityContentSelectorService(c)
	l := NewSecurityLdapService(c)
	s := NewSecuritySSLService(c)
	u := NewSecurityUserService(c)
	return &SecurityService{
		client: c,

		Anonymous:       a,
		ContentSelector: cs,
		LDAP:            l,
		SSL:             s,
		User:            u,
	}
}
