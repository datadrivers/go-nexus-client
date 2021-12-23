package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securitySamlAPIEndpoint = securityAPIEndpoint + "/saml"
)

type SecuritySamlService service

func NewSecuritySamlService(c *client) *SecuritySamlService {

	s := &SecuritySamlService{
		client: c,
	}
	return s
}

func (s *SecuritySamlService) Apply(saml security.SAML) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(saml)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(securitySamlAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if !(resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusNoContent) {
		return fmt.Errorf("could not create/update SAML configuration: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecuritySamlService) Read() (*security.SAML, error) {
	body, resp, err := s.client.Get(securitySamlAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get SAML configuration: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	samlServer := &security.SAML{}
	if err := json.Unmarshal(body, samlServer); err != nil {
		return nil, fmt.Errorf("could not unmarshal SAML configuration: %v", err)
	}

	return samlServer, nil
}

func (s *SecuritySamlService) Delete() error {
	body, resp, err := s.client.Delete(securitySamlAPIEndpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete SAML configuration: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}
