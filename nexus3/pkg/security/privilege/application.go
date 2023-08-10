package privilege

import (
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securityPrivilegesApplicationAPIEndpoint = securityPrivilegesAPIEndpoint + "/application"
)

type SecurityPrivilegeApplicationService struct {
	client *client.Client

	// Script *SecurityPrivilegeApplicationService
}

func NewSecurityPrivilegeApplicationService(c *client.Client) *SecurityPrivilegeApplicationService {
	return &SecurityPrivilegeApplicationService{
		client: c,
	}
}

func (s *SecurityPrivilegeApplicationService) Create(p security.PrivilegeApplication) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(securityPrivilegesApplicationAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create privilege \"%s\": HTTP: %d, %s", p.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityPrivilegeApplicationService) Update(name string, p security.PrivilegeApplication) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", securityPrivilegesApplicationAPIEndpoint, p.Name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update application privilege \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}
