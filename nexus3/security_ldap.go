package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
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

func (s *SecurityLdapService) List() ([]security.LDAP, error) {
	body, resp, err := s.client.Get(securityLdapAPIEndpoint, nil)
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

func (s *SecurityLdapService) Get(name string) (*security.LDAP, error) {
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", securityLdapAPIEndpoint, name), nil)
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
