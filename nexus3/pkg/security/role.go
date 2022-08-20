package security

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securityrolesAPIEndpoint = securityAPIEndpoint + "/roles"
)

type SecurityRoleService client.Service

func NewSecurityRoleService(c *client.Client) *SecurityRoleService {

	s := &SecurityRoleService{
		Client: c,
	}
	return s
}

func roleIOReader(role security.Role) (io.Reader, error) {
	b, err := json.Marshal(role)
	if err != nil {
		return nil, fmt.Errorf("could not marshal role data: %v", err)
	}

	return bytes.NewReader(b), nil
}

func (s *SecurityRoleService) Create(role security.Role) error {
	ioReader, err := roleIOReader(role)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Post(securityrolesAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *SecurityRoleService) Get(id string) (*security.Role, error) {
	encodedID := url.PathEscape(id)

	body, resp, err := s.Client.Get(fmt.Sprintf("%s/%s", securityrolesAPIEndpoint, encodedID), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}

	var role security.Role
	if err := json.Unmarshal(body, &role); err != nil {
		return nil, fmt.Errorf("could not unmarshal roles: %v", err)
	}
	return &role, nil

}

func (s *SecurityRoleService) Update(id string, role security.Role) error {
	encodedID := url.PathEscape(id)

	ioReader, err := roleIOReader(role)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/%s", securityrolesAPIEndpoint, encodedID), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *SecurityRoleService) Delete(id string) error {
	encodedID := url.PathEscape(id)

	body, resp, err := s.Client.Delete(fmt.Sprintf("%s/%s", securityrolesAPIEndpoint, encodedID))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}
