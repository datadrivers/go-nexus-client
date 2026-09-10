package license

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/license"
)

const (
	licenseAPIEndpoint = client.BasePath + "v1/system/license"
)

// ErrNoLicense is returned by Read when no license is installed on the Nexus
// instance.
var ErrNoLicense = errors.New("no license installed")

// LicenseService handles communication with the product licensing related
// methods
type LicenseService client.Service

// NewLicenseService creates a new instance of LicenseService
func NewLicenseService(c *client.Client) *LicenseService {

	s := &LicenseService{
		Client: c,
	}
	return s
}

// Read returns the details of the installed license. It returns ErrNoLicense
// when no license is installed.
func (s *LicenseService) Read() (*license.LicenseDetails, error) {
	body, resp, err := s.Client.Get(licenseAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNoLicense
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read license: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var licenseDetails license.LicenseDetails
	if err := json.Unmarshal(body, &licenseDetails); err != nil {
		return nil, fmt.Errorf("could not unmarshal license: %v", err)
	}

	return &licenseDetails, nil
}

// Install uploads the given license file content. The Nexus server must be
// restarted for the new license to take effect.
func (s *LicenseService) Install(licenseData []byte) (*license.LicenseDetails, error) {
	body, resp, err := s.Client.PostWithContentType(licenseAPIEndpoint, bytes.NewReader(licenseData), client.ContentTypeApplicationOctetStream)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("could not install license: HTTP %d, %s", resp.StatusCode, string(body))
	}

	if resp.StatusCode == http.StatusNoContent || len(body) == 0 {
		return nil, nil
	}

	var licenseDetails license.LicenseDetails
	if err := json.Unmarshal(body, &licenseDetails); err != nil {
		return nil, fmt.Errorf("could not unmarshal license: %v", err)
	}

	return &licenseDetails, nil
}

// Uninstall removes the installed license, if any
func (s *LicenseService) Uninstall() error {
	body, resp, err := s.Client.Delete(licenseAPIEndpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusNotFound {
		return fmt.Errorf("could not uninstall license: HTTP %d, %s", resp.StatusCode, string(body))
	}

	return nil
}
