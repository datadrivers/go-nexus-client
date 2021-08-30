package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	securityAnonymousAPIEndpoint = securityAPIEndpoint + "/anonymous"
)

type SecurityAnonymousService service

func NewSecurityAnonymousService(c *client) *SecurityAnonymousService {

	s := &SecurityAnonymousService{
		client: c,
	}
	return s
}

// Anonymous config
type SecurityAnonymousAccessSettings struct {
	// Whether or not Anonymous Access is enabled
	Enabled bool `json:"enabled"`

	// The username of the anonymous account
	UserID string `json:"userId"`

	// The name of the authentication realm for the anonymous account
	RealmName string `json:"realmName"`
}

// Get Anonymous Access settings
func (s *SecurityAnonymousService) Read() (*SecurityAnonymousAccessSettings, error) {
	body, resp, err := s.client.Get(securityAnonymousAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read anonymous config: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var anonymous SecurityAnonymousAccessSettings
	if err := json.Unmarshal(body, &anonymous); err != nil {
		return nil, fmt.Errorf("could not unmarshal anonymous config: %v", err)
	}

	return &anonymous, nil
}

func (s *SecurityAnonymousService) Update(anonymous SecurityAnonymousAccessSettings) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(anonymous)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(securityAnonymousAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not update anonymous config: HTTP %d, %s", resp.StatusCode, string(body))
	}

	return nil
}
