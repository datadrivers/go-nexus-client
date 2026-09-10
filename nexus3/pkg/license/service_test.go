package license

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/stretchr/testify/assert"
)

var (
	testClient *client.Client = nil
)

func getTestClient() *client.Client {
	if testClient != nil {
		return testClient
	}
	return client.NewClient(getDefaultConfig())
}

func getTestService() *LicenseService {
	return NewLicenseService(getTestClient())
}

func getDefaultConfig() client.Config {
	return client.Config{
		Insecure: tools.GetEnv("NEXUS_INSECURE_SKIP_VERIFY", true).(bool),
		Password: tools.GetEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      tools.GetEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: tools.GetEnv("NEXUS_USRNAME", "admin").(string),
	}
}

func TestNewLicenseService(t *testing.T) {
	s := getTestService()

	assert.NotNil(t, s, "NewLicenseService() must not return nil")
}

// TestLicenseRead only asserts that reading the license succeeds or reports
// that none is installed. The test instance is a Community Edition one, so no
// license can be installed or uninstalled against it.
func TestLicenseRead(t *testing.T) {
	service := getTestService()

	licenseDetails, err := service.Read()
	if err == ErrNoLicense {
		assert.Nil(t, licenseDetails)
		return
	}

	assert.Nil(t, err)
	assert.NotNil(t, licenseDetails)
}
