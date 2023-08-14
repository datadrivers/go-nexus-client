package deprecated

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securityPrivilegesAPIEndpoint = securityAPIEndpoint + "/privileges"
)

type SecurityPrivilegeService client.Service

func NewSecurityPrivilegeService(c *client.Client) *SecurityPrivilegeService {

	s := &SecurityPrivilegeService{
		Client: c,
	}
	return s
}

func (s *SecurityPrivilegeService) List() ([]security.Privilege, error) {
	body, resp, err := s.Client.Get(securityPrivilegesAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read privileges: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var privileges []security.Privilege
	if err := json.Unmarshal(body, &privileges); err != nil {
		return nil, fmt.Errorf("could not unmarshal privileges: %v", err)
	}

	return privileges, nil
}

func (s *SecurityPrivilegeService) Create(p security.Privilege) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Post(fmt.Sprintf("%s/%s", securityPrivilegesAPIEndpoint, strings.ToLower(p.Type)), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create privilege \"%s\": HTTP: %d, %s", p.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityPrivilegeService) Get(name string) (*security.Privilege, error) {
	privileges, err := s.List()
	if err != nil {
		return nil, err
	}

	for _, p := range privileges {
		if p.Name == name {
			return &p, nil
		}
	}

	return nil, nil
}

func (s *SecurityPrivilegeService) Update(name string, p security.Privilege) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/%s/%s", securityPrivilegesAPIEndpoint, p.Type, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update privilege \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityPrivilegeService) Delete(name string) error {
	body, resp, err := s.Client.Delete(fmt.Sprintf("%s/%s", securityPrivilegesAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete privilege \"%s\": HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	return nil
}
