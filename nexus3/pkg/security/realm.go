package security

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/security"
)

const (
	securityRealmAPIEndpoint = securityAPIEndpoint + "/realms"
)

type SecurityRealmService client.Service

func NewSecurityRealmService(c *client.Client) *SecurityRealmService {

	s := &SecurityRealmService{
		Client: c,
	}
	return s
}

func (s *SecurityRealmService) Activate(ids []string) error {
	data, err := json.Marshal(ids)
	if err != nil {
		return fmt.Errorf("could not marshal realm IDs: %v", err)
	}

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/active", securityRealmAPIEndpoint), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("could not activate realms: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not activate realms: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityRealmService) ListActive() ([]string, error) {
	body, resp, err := s.Client.Get(fmt.Sprintf("%s/active", securityRealmAPIEndpoint), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read active realms: HTTP: %d, %v", resp.StatusCode, err)
	}

	var result []string
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("could not unmarshal active realms: %v", err)
	}
	return result, nil
}

func (s *SecurityRealmService) ListAvailable() ([]security.Realm, error) {
	body, resp, err := s.Client.Get(fmt.Sprintf("%s/available", securityRealmAPIEndpoint), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read available realms: HTTP: %d, %v", resp.StatusCode, err)
	}

	var result []security.Realm
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("could not unmarshal available realms: %v", err)
	}
	return result, nil
}
