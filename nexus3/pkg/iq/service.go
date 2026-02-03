package iq

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/tools"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/iq"
)

const (
	iqAPIEndpoint = "service/rest/v1/iq"
)

// IQServerService handles communication with the IQ Server configuration related methods
type IQServerService struct {
	client *client.Client
}

// NewIQServerService creates a new instance of IQServerService
func NewIQServerService(c *client.Client) *IQServerService {
	return &IQServerService{
		client: c,
	}
}

func jsonUnmarshalIQServerConfig(data []byte) (*iq.IQServerConfiguration, error) {
	var config iq.IQServerConfiguration
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("could not unmarshal IQ Server configuration: %w", err)
	}
	return &config, nil
}

// Get retrieves the current IQ Server configuration
func (s *IQServerService) Get() (*iq.IQServerConfiguration, error) {
	body, resp, err := s.client.Get(iqAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get IQ Server configuration: HTTP %d", resp.StatusCode)
	}

	return jsonUnmarshalIQServerConfig(body)
}

// Update updates the IQ Server configuration
func (s *IQServerService) Update(config iq.IQServerConfiguration) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(config)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(iqAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not update IQ Server configuration: HTTP %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// Enable enables the IQ Server connection
func (s *IQServerService) Enable() error {
	config, err := s.Get()
	if err != nil {
		return err
	}

	config.Enabled = true
	return s.Update(*config)
}

// Disable disables the IQ Server connection
func (s *IQServerService) Disable() error {
	config, err := s.Get()
	if err != nil {
		return err
	}

	config.Enabled = false
	return s.Update(*config)
}

// VerifyConnection verifies the connection to the configured IQ Server
// This endpoint is available in Nexus Repository Manager 3.24+
func (s *IQServerService) VerifyConnection() error {
	_, resp, err := s.client.Post(iqAPIEndpoint+"/verify-connection", nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not verify IQ Server connection: HTTP %d", resp.StatusCode)
	}

	return nil
}
