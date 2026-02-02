package capability

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/capability"
)

const (
	capabilitiesAPIEndpoint = client.BasePath + "v1/capabilities"
)

type CapabilityService client.Service

func NewCapabilityService(c *client.Client) *CapabilityService {
	s := &CapabilityService{
		Client: c,
	}
	return s
}

func jsonUnmarshalCapabilities(data []byte) ([]capability.Capability, error) {
	var capabilities []capability.Capability
	if err := json.Unmarshal(data, &capabilities); err != nil {
		return nil, fmt.Errorf("could not unmarshal capabilities: %v", err)
	}
	return capabilities, nil
}

func jsonUnmarshalCapability(data []byte) (*capability.Capability, error) {
	var cap capability.Capability
	if err := json.Unmarshal(data, &cap); err != nil {
		return nil, fmt.Errorf("could not unmarshal capability: %v", err)
	}
	return &cap, nil
}

// List returns all capabilities
func (s *CapabilityService) List() ([]capability.Capability, error) {
	body, resp, err := s.Client.Get(capabilitiesAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not list capabilities: HTTP %d: %s", resp.StatusCode, string(body))
	}

	return jsonUnmarshalCapabilities(body)
}

// Get returns a specific capability by ID
// Note: Nexus API doesn't support GET by ID, so we list all and filter
func (s *CapabilityService) Get(id string) (*capability.Capability, error) {
	capabilities, err := s.List()
	if err != nil {
		return nil, err
	}

	for _, cap := range capabilities {
		if cap.ID == id {
			return &cap, nil
		}
	}

	return nil, nil // Not found
}

// Create creates a new capability
func (s *CapabilityService) Create(cap capability.CapabilityCreate) (*capability.Capability, error) {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(cap)
	if err != nil {
		return nil, err
	}

	body, resp, err := s.Client.Post(capabilitiesAPIEndpoint, ioReader)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("could not create capability: HTTP %d: %s", resp.StatusCode, string(body))
	}

	return jsonUnmarshalCapability(body)
}

// Update updates an existing capability
func (s *CapabilityService) Update(id string, cap capability.CapabilityUpdate) error {
	endpoint := fmt.Sprintf("%s/%s", capabilitiesAPIEndpoint, id)

	ioReader, err := tools.JsonMarshalInterfaceToIOReader(cap)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(endpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not update capability: HTTP %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// Delete deletes a capability by ID
func (s *CapabilityService) Delete(id string) error {
	endpoint := fmt.Sprintf("%s/%s", capabilitiesAPIEndpoint, id)

	body, resp, err := s.Client.Delete(endpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not delete capability: HTTP %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// ListTypes returns all available capability types
func (s *CapabilityService) ListTypes() ([]capability.TypeDescriptor, error) {
	endpoint := fmt.Sprintf("%s/types", capabilitiesAPIEndpoint)

	body, resp, err := s.Client.Get(endpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not list capability types: HTTP %d: %s", resp.StatusCode, string(body))
	}

	var types []capability.TypeDescriptor
	if err := json.Unmarshal(body, &types); err != nil {
		return nil, fmt.Errorf("could not unmarshal capability types: %v", err)
	}

	return types, nil
}
