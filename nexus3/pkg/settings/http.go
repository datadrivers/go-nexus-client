package settings

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/settings"
)

const (
	settingsHTTPAPIEndpoint = settingsAPIEndpoint + "/http"
)

type SettingsHTTPService client.Service

func NewSettingsHTTPService(c *client.Client) *SettingsHTTPService {

	s := &SettingsHTTPService{
		Client: c,
	}
	return s
}

// Read returns the outbound HTTP system settings
func (s *SettingsHTTPService) Read() (*settings.HTTPSettings, error) {
	body, resp, err := s.Client.Get(settingsHTTPAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read http settings: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var httpSettings settings.HTTPSettings
	if err := json.Unmarshal(body, &httpSettings); err != nil {
		return nil, fmt.Errorf("could not unmarshal http settings: %v", err)
	}

	return &httpSettings, nil
}

// Update applies the given outbound HTTP system settings
func (s *SettingsHTTPService) Update(httpSettings settings.HTTPSettings) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(httpSettings)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(settingsHTTPAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not update http settings: HTTP %d, %s", resp.StatusCode, string(body))
	}

	return nil
}

// Reset restores the outbound HTTP system settings to their defaults
func (s *SettingsHTTPService) Reset() error {
	body, resp, err := s.Client.Delete(settingsHTTPAPIEndpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not reset http settings: HTTP %d, %s", resp.StatusCode, string(body))
	}

	return nil
}
