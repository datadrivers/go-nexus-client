package privilege

import (
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securityWildcardAPIEndpoint = securityPrivilegesAPIEndpoint + "/wildcard"
)

type SecurityPrivilegeWildcardService struct {
	client *client.Client
}

func NewSecurityPrivilegeWildcardService(c *client.Client) *SecurityPrivilegeWildcardService {
	return &SecurityPrivilegeWildcardService{
		client: c,
	}
}

func (s *SecurityPrivilegeWildcardService) Create(p security.PrivilegeWildcard) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(securityWildcardAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create wildcard privilege \"%s\": HTTP: %d, %s", p.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityPrivilegeWildcardService) Update(name string, p security.PrivilegeWildcard) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", securityWildcardAPIEndpoint, p.Name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update wildcard privilege \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}
