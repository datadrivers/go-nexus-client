package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securityUserTokensAPIEndpoint = securityAPIEndpoint + "/user-tokens"
)

type SecurityUserTokensService service

func NewSecurityUserTokensService(c *client) *SecurityUserTokensService {

	s := &SecurityUserTokensService{
		client: c,
	}
	return s
}

func (s *SecurityUserTokensService) Configure(userTokens security.UserTokenConfiguration) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(userTokens)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(securityUserTokensAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if !(resp.StatusCode == http.StatusOK) {
		return fmt.Errorf("could not create/update UserTokenConfiguration configuration: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityUserTokensService) Get() (*security.UserTokenConfiguration, error) {
	body, resp, err := s.client.Get(securityUserTokensAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get UserTokenConfiguration configuration: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	userTokensServer := &security.UserTokenConfiguration{}
	if err := json.Unmarshal(body, userTokensServer); err != nil {
		return nil, fmt.Errorf("could not unmarshal UserTokenConfiguration configuration: %v", err)
	}

	return userTokensServer, nil
}
