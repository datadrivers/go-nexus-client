package security

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/tools"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/security"
)

const (
	securityLdapAPIEndpoint = securityAPIEndpoint + "/ldap"
)

type SecurityLdapService client.Service

func NewSecurityLdapService(c *client.Client) *SecurityLdapService {

	s := &SecurityLdapService{
		Client: c,
	}
	return s
}

func (s *SecurityLdapService) ChangeOrder(order []string) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(order)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Post(fmt.Sprintf("%s/change-order", securityLdapAPIEndpoint), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not change LDAP order: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityLdapService) List() ([]security.LDAP, error) {
	body, resp, err := s.Client.Get(securityLdapAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get LDAP server: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	var result []security.LDAP
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("could not unmarshal LDAP server: %v", err)
	}

	return result, nil
}

func (s *SecurityLdapService) Create(ldap security.LDAP) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(ldap)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Post(securityLdapAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create LDAP server: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityLdapService) Get(name string) (*security.LDAP, error) {
	body, resp, err := s.Client.Get(fmt.Sprintf("%s/%s", securityLdapAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get LDAP server '%s': HTTP: %d, %v", name, resp.StatusCode, string(body))
	}

	ldapServer := &security.LDAP{}
	if err := json.Unmarshal(body, ldapServer); err != nil {
		return nil, fmt.Errorf("could not unmarshal LDAP server '%s': %v", name, err)
	}

	return ldapServer, nil
}

func (s *SecurityLdapService) Update(name string, ldap security.LDAP) error {
	if ldap.ID == "" {
		ldapFound, err := s.Get(ldap.Name)
		if err != nil {
			return err
		}
		ldap.ID = ldapFound.ID

	}
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(ldap)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/%s", securityLdapAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update LDAP server `%s`: HTTP: %d, :%v", name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityLdapService) Delete(name string) error {
	body, resp, err := s.Client.Delete(fmt.Sprintf("%s/%s", securityLdapAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete LDAP server '%s': HTTP: %d, %v", name, resp.StatusCode, string(body))
	}

	return nil
}
