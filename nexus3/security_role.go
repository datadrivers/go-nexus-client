package nexus3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securityrolesAPIEndpoint = securityAPIEndpoint + "/roles"
)

type SecurityRoleService service

func NewSecurityRoleService(c *client) *SecurityRoleService {

	s := &SecurityRoleService{
		client: c,
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

	body, resp, err := s.client.Post(securityrolesAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *SecurityRoleService) Get(id string) (*security.Role, error) {
	body, resp, err := s.client.Get(securityrolesAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}

	var roles []security.Role
	if err := json.Unmarshal(body, &roles); err != nil {
		return nil, fmt.Errorf("could not unmarshal roles: %v", err)
	}

	for _, role := range roles {
		if role.ID == id {
			return &role, nil
		}
	}

	return nil, nil
}

func (s *SecurityRoleService) Update(id string, role security.Role) error {
	ioReader, err := roleIOReader(role)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", securityrolesAPIEndpoint, id), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *SecurityRoleService) Delete(id string) error {
	body, resp, err := s.client.Delete(fmt.Sprintf("%s/%s", securityrolesAPIEndpoint, id))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}
