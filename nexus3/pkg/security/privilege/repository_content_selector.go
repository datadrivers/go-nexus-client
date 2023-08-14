package privilege

import (
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securityContentSelectorAPIEndpoint = securityPrivilegesAPIEndpoint + "/repository-content-selector"
)

type SecurityPrivilegeContentSelectorService struct {
	client *client.Client
}

func NewSecurityPrivilegeContentSelectorService(c *client.Client) *SecurityPrivilegeContentSelectorService {
	return &SecurityPrivilegeContentSelectorService{
		client: c,
	}
}

func (s *SecurityPrivilegeContentSelectorService) Create(p security.PrivilegeRepositoryContentSelector) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(securityContentSelectorAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create privilege \"%s\": HTTP: %d, %s", p.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s *SecurityPrivilegeContentSelectorService) Update(name string, p security.PrivilegeRepositoryContentSelector) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(p)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", securityContentSelectorAPIEndpoint, p.Name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update privilege \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}
