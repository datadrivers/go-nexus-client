package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ldapAPIEndpoint = basePath + "v1/security/ldap"
)

// LDAP data structure
type LDAP struct {
	AuthPassword                string `json:"authPassword"`
	AuthRealm                   string `json:"authRealm,omitempty"`
	AuthSchema                  string `json:"authScheme"`
	AuthUserName                string `json:"authUsername,omitempty"`
	ConnectionRetryDelaySeconds uint   `json:"connectionRetryDelaySeconds"`
	ConnectionTimeoutSeconds    uint   `json:"connectionTimeoutSeconds"`
	GroupBaseDn                 string `json:"groupBaseDn,omitempty"`
	GroupIDAttribute            string `json:"groupIdAttribute,omitempty"`
	GroupMemberAttribute        string `json:"groupMemberAttribute,omitempty"`
	GroupMemberFormat           string `json:"groupMemberFormat,omitempty"`
	GroupObjectClass            string `json:"groupObjectClass,omitempty"`
	GroupSubtree                bool   `json:"groupSubtree,omitempty"`
	GroupType                   string `json:"groupType"`
	Host                        string `json:"host"`
	ID                          string `json:"id"`
	LDAPGroupsAsRoles           bool   `json:"ldapGroupsAsRoles,omitempty"`
	MaxIncidentCount            uint   `json:"maxIncidentsCount"`
	Name                        string `json:"name"`
	Port                        uint   `json:"port"`
	Protocol                    string `json:"protocol"`
	SearchBase                  string `json:"searchBase"`
	UseTrustStore               bool   `json:"useTrustStore,omitempty"`
	UserBaseDN                  string `json:"userBaseDn,omitempty"`
	UserEmailAddressAttribute   string `json:"userEmailAddressAttribute,omitempty"`
	UserIDAttribute             string `json:"userIdAttribute,omitempty"`
	UserLDAPFilter              string `json:"userLdapFilter,omitempty"`
	UserMemberOfAttribute       string `json:"userMemberOfAttribute,omitempty"`
	UserObjectClass             string `json:"userObjectClass,omitempty"`
	UserPasswordAttribute       string `json:"userPasswordAttribute,omitempty"`
	UserRealNameAttribute       string `json:"userRealNameAttribute,omitempty"`
	UserSubtree                 bool   `json:"userSubtree,omitempty"`
}

func (c *client) LDAPChangeOrder(order []string) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(order)
	if err != nil {
		return err
	}

	body, resp, err := c.Post(fmt.Sprintf("%s/change-order", ldapAPIEndpoint), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not change LDAP order: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (c *client) LDAPList() ([]LDAP, error) {
	body, resp, err := c.Get(ldapAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get LDAP server: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	var result []LDAP
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("could not unmarshal LDAP server: %v", err)
	}

	return result, nil
}

func (c *client) LDAPCreate(ldap LDAP) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(ldap)
	if err != nil {
		return err
	}

	body, resp, err := c.Post(ldapAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create LDAP server: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (c *client) LDAPRead(name string) (*LDAP, error) {
	body, resp, err := c.Get(fmt.Sprintf("%s/%s", ldapAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get LDAP server '%s': HTTP: %d, %v", name, resp.StatusCode, string(body))
	}

	ldapServer := &LDAP{}
	if err := json.Unmarshal(body, ldapServer); err != nil {
		return nil, fmt.Errorf("could not unmarshal LDAP server '%s': %v", name, err)
	}

	return ldapServer, nil
}

func (c *client) LDAPUpdate(name string, ldap LDAP) error {
	if ldap.ID == "" {
		ldapFound, err := c.LDAPRead(ldap.Name)
		if err != nil {
			return err
		}
		ldap.ID = ldapFound.ID

	}
	ioReader, err := jsonMarshalInterfaceToIOReader(ldap)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(fmt.Sprintf("%s/%s", ldapAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update LDAP server `%s`: HTTP: %d, :%v", name, resp.StatusCode, string(body))
	}

	return nil
}

func (c *client) LDAPDelete(name string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", ldapAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete LDAP server '%s': HTTP: %d, %v", name, resp.StatusCode, string(body))
	}

	return nil
}
