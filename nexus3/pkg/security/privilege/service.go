package privilege

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securityPrivilegesAPIEndpoint = client.BasePath + "v1/security/privileges"
)

type SecurityPrivilegeService struct {
	client *client.Client

	// API Services
	Script                    *SecurityPrivilegeScriptService
	Application               *SecurityPrivilegeApplicationService
	RepositoryAdmin           *SecurityPrivilegeRepositoryAdminService
	RepositoryContentSelector *SecurityPrivilegeContentSelectorService
	RepositoryView            *SecurityPrivilegeRepositoryViewService
	Wildcard                  *SecurityPrivilegeWildcardService
}

func NewSecurityPrivilegeService(c *client.Client) *SecurityPrivilegeService {
	return &SecurityPrivilegeService{
		client: c,

		Script:                    NewSecurityPrivilegeScriptService(c),
		Application:               NewSecurityPrivilegeApplicationService(c),
		RepositoryAdmin:           NewSecurityPrivilegeRepositoryAdminService(c),
		RepositoryContentSelector: NewSecurityPrivilegeContentSelectorService(c),
		RepositoryView:            NewSecurityPrivilegeRepositoryViewService(c),
		Wildcard:                  NewSecurityPrivilegeWildcardService(c),
	}
}

func listPrivileges(c *client.Client) ([]security.Privilege, error) {
	body, resp, err := c.Get(securityPrivilegesAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not list privileges: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var privileges []security.Privilege
	if err := json.Unmarshal(body, &privileges); err != nil {
		return nil, fmt.Errorf("could not unmarshal list of privileges: %v", err)
	}

	return privileges, nil
}

func (s SecurityPrivilegeService) List() ([]security.Privilege, error) {
	return listPrivileges(s.client)
}

func (s SecurityPrivilegeService) Get(name string) (*security.Privilege, error) {
	var privilege security.Privilege
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", securityPrivilegesAPIEndpoint, name), nil)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read privilege '%s': HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	if err := json.Unmarshal(body, &privilege); err != nil {
		return nil, fmt.Errorf("could not unmarshal privilege: %v", err)
	}
	return &privilege, nil
}

func (s SecurityPrivilegeService) Delete(name string) error {
	body, resp, err := s.client.Delete(fmt.Sprintf("%s/%s", securityPrivilegesAPIEndpoint, name))

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}
	return err
}
