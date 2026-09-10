package security

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func TestSecuritySsrfProtection(t *testing.T) {
	service := getTestService()

	oldSsrfProtection, err := service.SsrfProtection.Read()
	assert.Nil(t, err)
	assert.NotNil(t, oldSsrfProtection)

	newSsrfProtection := security.SsrfProtectionConfiguration{
		Enabled:        true,
		AllowedIPs:     []string{"10.0.0.50", "192.168.1.100"},
		AllowedDomains: []string{"internal.corp.com", "registry.local"},
	}
	updated, err := service.SsrfProtection.Update(newSsrfProtection)
	assert.Nil(t, err)
	assert.NotNil(t, updated)

	ssrfProtection, err := service.SsrfProtection.Read()
	assert.Nil(t, err)
	assert.NotNil(t, ssrfProtection)
	assert.Equal(t, newSsrfProtection.Enabled, ssrfProtection.Enabled)
	assert.ElementsMatch(t, newSsrfProtection.AllowedIPs, ssrfProtection.AllowedIPs)
	assert.ElementsMatch(t, newSsrfProtection.AllowedDomains, ssrfProtection.AllowedDomains)

	_, err = service.SsrfProtection.Update(*oldSsrfProtection)
	assert.Nil(t, err)
}
