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
	securityAnonymousAPIEndpoint = securityAPIEndpoint + "/anonymous"
)

type SecurityAnonymousService client.Service

func NewSecurityAnonymousService(c *client.Client) *SecurityAnonymousService {

	s := &SecurityAnonymousService{
		Client: c,
	}
	return s
}

// Get Anonymous Access settings
func (s *SecurityAnonymousService) Read() (*security.AnonymousAccessSettings, error) {
	body, resp, err := s.Client.Get(securityAnonymousAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read anonymous config: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var anonymous security.AnonymousAccessSettings
	if err := json.Unmarshal(body, &anonymous); err != nil {
		return nil, fmt.Errorf("could not unmarshal anonymous config: %v", err)
	}

	return &anonymous, nil
}

func (s *SecurityAnonymousService) Update(anonymous security.AnonymousAccessSettings) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(anonymous)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(securityAnonymousAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not update anonymous config: HTTP %d, %s", resp.StatusCode, string(body))
	}

	return nil
}
