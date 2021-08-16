package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	samlAPIEndpoint = basePath + "v1/security/saml"
)

// SAML data structure
type SAML struct {
	EntityId                   string `json:"entityId,omitempty"`
	IdpMetadata                string `json:"idpMetadata"`
	UsernameAttribute          string `json:"usernameAttribute,omitempty"`
	FirstNameAttribute         string `json:"firstNameAttribute,omitempty"`
	LastNameAttribute          string `json:"lastNameAttribute,omitempty"`
	EmailAttribute             string `json:"emailAttribute,omitempty"`
	GroupsAttribute            string `json:"groupsAttribute,omitempty"`
	ValidateResponseSignature  bool   `json:"validateResponseSignature,omitempty"`
	ValidateAssertionSignature bool   `json:"validateAssertionSignature,omitempty"`
}

func (c *client) SAMLApply(saml SAML) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(saml)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(samlAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if !(resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusNoContent) {
		return fmt.Errorf("could not create/update SAML configuration: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (c *client) SAMLRead() (*SAML, error) {
	body, resp, err := c.Get(samlAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get SAML configuration: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	samlServer := &SAML{}
	if err := json.Unmarshal(body, samlServer); err != nil {
		return nil, fmt.Errorf("could not unmarshal SAML configuration: %v", err)
	}

	return samlServer, nil
}

func (c *client) SAMLDelete() error {
	body, resp, err := c.Delete(samlAPIEndpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete SAML configuration: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}
