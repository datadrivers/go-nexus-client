package privilege

import (
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securityRepositoryViewAPIEndpoint = securityPrivilegesAPIEndpoint + "/repository-view"
)

type SecurityPrivilegeRepositoryViewService struct {
	client *client.Client
}

func NewSecurityPrivilegeRepositoryViewService(c *client.Client) *SecurityPrivilegeRepositoryViewService {
	return &SecurityPrivilegeRepositoryViewService{
		client: c,
	}
}

func (s *SecurityPrivilegeRepositoryViewService) Create(p security.PrivilegeRepositoryView) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(securityRepositoryViewAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create privilege \"%s\": HTTP: %d, %s", p.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityPrivilegeRepositoryViewService) Update(name string, p security.PrivilegeRepositoryView) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", securityRepositoryViewAPIEndpoint, p.Name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update privilege \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}
