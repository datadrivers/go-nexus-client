package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	securityLdapAPIEndpoint = securityAPIEndpoint + "/ldap"
)

type SecurityLdapService service

func NewSecurityLdapService(c *client) *SecurityLdapService {

	s := &SecurityLdapService{
		client: c,
	}
	return s
}

// SecurityLDAP data structure
type SecurityLDAP struct {
	// The password to bind with. Required if authScheme other than none
	AuthPassword string `json:"authPassword"`

	// The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5
	AuthRealm string `json:"authRealm,omitempty"`

	// Authentication scheme used for connecting to LDAP server
	AuthSchema string `json:"authScheme"`

	// This must be a fully qualified username if simple authentication is used. Required if authScheme other than none
	AuthUserName string `json:"authUsername,omitempty"`

	// How long to wait before retrying
	ConnectionRetryDelaySeconds int32 `json:"connectionRetryDelaySeconds"`

	// How long to wait before timeout
	ConnectionTimeoutSeconds int32 `json:"connectionTimeoutSeconds"`

	// The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN
	GroupBaseDn string `json:"groupBaseDn,omitempty"`

	// This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static
	GroupIDAttribute string `json:"groupIdAttribute,omitempty"`

	// LDAP attribute containing the usernames for the group. Required if groupType is static
	GroupMemberAttribute string `json:"groupMemberAttribute,omitempty"`

	// The format of user ID stored in the group member attribute. Required if groupType is static
	GroupMemberFormat string `json:"groupMemberFormat,omitempty"`

	// LDAP class for group objects. Required if groupType is static
	GroupObjectClass string `json:"groupObjectClass,omitempty"`

	// Are groups located in structures below the group base DN
	GroupSubtree bool `json:"groupSubtree,omitempty"`

	// Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true
	GroupType string `json:"groupType"`

	// LDAP server connection hostname
	Host string `json:"host"`
	ID   string `json:"id"`

	// Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles
	LDAPGroupsAsRoles bool `json:"ldapGroupsAsRoles,omitempty"`

	// How many retry attempts
	MaxIncidentCount int32 `json:"maxIncidentsCount"`

	// LDAP server name
	Name string `json:"name"`

	// LDAP server connection port to use
	Port int32 `json:"port"`

	// LDAP server connection Protocol to use
	Protocol string `json:"protocol"`

	// LDAP location to be added to the connection URL
	SearchBase string `json:"searchBase"`

	// Whether to use certificates stored in Nexus Repository Manager's truststore
	UseTrustStore bool `json:"useTrustStore,omitempty"`

	// The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN
	UserBaseDN string `json:"userBaseDn,omitempty"`

	// This is used to find an email address given the user ID
	UserEmailAddressAttribute string `json:"userEmailAddressAttribute,omitempty"`

	// This is used to find a user given its user ID
	UserIDAttribute string `json:"userIdAttribute,omitempty"`

	// LDAP search filter to limit user search. example: "(|(mail=*@example.com)(uid=dom*))"
	UserLDAPFilter string `json:"userLdapFilter,omitempty"`

	// Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic
	UserMemberOfAttribute string `json:"userMemberOfAttribute,omitempty"`

	// LDAP class for user objects
	UserObjectClass string `json:"userObjectClass,omitempty"`

	// If this field is blank the user will be authenticated against a bind with the LDAP server
	UserPasswordAttribute string `json:"userPasswordAttribute,omitempty"`

	// This is used to find a real name given the user ID
	UserRealNameAttribute string `json:"userRealNameAttribute,omitempty"`

	// Are users located in structures below the user base DN?
	UserSubtree bool `json:"userSubtree,omitempty"`
}

func (s *SecurityLdapService) ChangeOrder(order []string) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(order)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(fmt.Sprintf("%s/change-order", securityLdapAPIEndpoint), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not change LDAP order: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityLdapService) List() ([]SecurityLDAP, error) {
	body, resp, err := s.client.Get(securityLdapAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get LDAP server: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	var result []SecurityLDAP
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("could not unmarshal LDAP server: %v", err)
	}

	return result, nil
}

func (s *SecurityLdapService) Create(ldap SecurityLDAP) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(ldap)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(securityLdapAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create LDAP server: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityLdapService) Get(name string) (*SecurityLDAP, error) {
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", securityLdapAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get LDAP server '%s': HTTP: %d, %v", name, resp.StatusCode, string(body))
	}

	ldapServer := &SecurityLDAP{}
	if err := json.Unmarshal(body, ldapServer); err != nil {
		return nil, fmt.Errorf("could not unmarshal LDAP server '%s': %v", name, err)
	}

	return ldapServer, nil
}

func (s *SecurityLdapService) Update(name string, ldap SecurityLDAP) error {
	if ldap.ID == "" {
		ldapFound, err := s.Get(ldap.Name)
		if err != nil {
			return err
		}
		ldap.ID = ldapFound.ID

	}
	ioReader, err := jsonMarshalInterfaceToIOReader(ldap)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", securityLdapAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update LDAP server `%s`: HTTP: %d, :%v", name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityLdapService) Delete(name string) error {
	body, resp, err := s.client.Delete(fmt.Sprintf("%s/%s", securityLdapAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete LDAP server '%s': HTTP: %d, %v", name, resp.StatusCode, string(body))
	}

	return nil
}
