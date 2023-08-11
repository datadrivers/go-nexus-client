package privilege

import (
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securityPrivilegesScriptAPIEndpoint = securityPrivilegesAPIEndpoint + "/script"
)

type SecurityPrivilegeScriptService struct {
	client *client.Client

	// Script *SecurityPrivilegeScriptService
}

func NewSecurityPrivilegeScriptService(c *client.Client) *SecurityPrivilegeScriptService {
	return &SecurityPrivilegeScriptService{
		client: c,
	}
}

func (s *SecurityPrivilegeScriptService) Create(p security.PrivilegeScript) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(securityPrivilegesScriptAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create privilege \"%s\": HTTP: %d, %s", p.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityPrivilegeScriptService) Update(name string, p security.PrivilegeScript) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", securityPrivilegesScriptAPIEndpoint, p.Name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update privilege \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}
