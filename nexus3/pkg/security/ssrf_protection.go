package security

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securitySsrfProtectionAPIEndpoint = securityAPIEndpoint + "/ssrf-protection"
)

type SecuritySsrfProtectionService client.Service

func NewSecuritySsrfProtectionService(c *client.Client) *SecuritySsrfProtectionService {

	s := &SecuritySsrfProtectionService{
		Client: c,
	}
	return s
}

// Read returns the SSRF protection settings
func (s *SecuritySsrfProtectionService) Read() (*security.SsrfProtectionConfiguration, error) {
	body, resp, err := s.Client.Get(securitySsrfProtectionAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read ssrf protection config: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var ssrfProtection security.SsrfProtectionConfiguration
	if err := json.Unmarshal(body, &ssrfProtection); err != nil {
		return nil, fmt.Errorf("could not unmarshal ssrf protection config: %v", err)
	}

	return &ssrfProtection, nil
}

// Update applies the given SSRF protection settings and returns the resulting
// configuration as reported by Nexus
func (s *SecuritySsrfProtectionService) Update(ssrfProtection security.SsrfProtectionConfiguration) (*security.SsrfProtectionConfiguration, error) {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(ssrfProtection)
	if err != nil {
		return nil, err
	}

	body, resp, err := s.Client.Put(securitySsrfProtectionAPIEndpoint, ioReader)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("could not update ssrf protection config: HTTP %d, %s", resp.StatusCode, string(body))
	}

	if resp.StatusCode == http.StatusNoContent || len(body) == 0 {
		return &ssrfProtection, nil
	}

	var updated security.SsrfProtectionConfiguration
	if err := json.Unmarshal(body, &updated); err != nil {
		return nil, fmt.Errorf("could not unmarshal ssrf protection config: %v", err)
	}

	return &updated, nil
}
