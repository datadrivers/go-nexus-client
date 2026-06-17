package iq

import (
	"encoding/json"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/iq"
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

func getTestService() *IQServerService {
	return NewIQServerService(getTestClient())
}

func getDefaultConfig() client.Config {
	insecure := true
	if val := tools.GetEnv("NEXUS_INSECURE_SKIP_VERIFY", "true"); val != nil {
		if strVal, ok := val.(string); ok {
			insecure = strVal == "true" || strVal == "1"
		} else if boolVal, ok := val.(bool); ok {
			insecure = boolVal
		}
	}

	return client.Config{
		Insecure: insecure,
		Password: tools.GetEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      tools.GetEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: tools.GetEnv("NEXUS_USERNAME", "admin").(string),
	}
}

func TestNewIQServerService(t *testing.T) {
	s := getTestService()
	assert.NotNil(t, s, "NewIQServerService() must not return nil")
}

func TestJSONUnmarshalIQServerConfig(t *testing.T) {
	url := "https://iq.example.com"
	authType := "USER"
	username := "admin"
	password := "password123"
	timeout := 60

	testConfig := iq.IQServerConfiguration{
		Enabled:             true,
		ShowLink:            true,
		URL:                 &url,
		AuthenticationType:  &authType,
		Username:            &username,
		Password:            &password,
		UseTrustStoreForURL: false,
		TimeoutSeconds:      &timeout,
		Properties:          nil,
		FailOpenModeEnabled: false,
	}

	testData, err := json.Marshal(testConfig)
	if err != nil {
		t.Fatalf("could not marshal testConfig: %v", err)
	}

	config, err := jsonUnmarshalIQServerConfig(testData)
	assert.Nil(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, true, config.Enabled)
	assert.Equal(t, url, *config.URL)
	assert.Equal(t, authType, *config.AuthenticationType)
	assert.Equal(t, username, *config.Username)
}

func TestIQServerGetConfiguration(t *testing.T) {
	service := getTestService()

	config, err := service.Get()
	assert.Nil(t, err)
	assert.NotNil(t, config)
	// The configuration should exist even if not enabled
	// Initial state should be disabled
}

func TestIQServerUpdateConfiguration(t *testing.T) {
	service := getTestService()

	// Get current configuration
	originalConfig, err := service.Get()
	assert.Nil(t, err)
	assert.NotNil(t, originalConfig)

	// Save original values to restore later
	originalEnabled := originalConfig.Enabled
	originalURL := originalConfig.URL
	originalAuthType := originalConfig.AuthenticationType
	originalUsername := originalConfig.Username

	// Test update with minimal valid configuration (disabled)
	url := "https://iq-test.example.com"
	authType := "USER"
	username := "testuser"
	password := "testpass"
	timeout := 30

	testConfig := iq.IQServerConfiguration{
		Enabled:             false, // Keep disabled to avoid connection issues
		ShowLink:            false,
		URL:                 &url,
		AuthenticationType:  &authType,
		Username:            &username,
		Password:            &password,
		UseTrustStoreForURL: false,
		TimeoutSeconds:      &timeout,
		Properties:          nil,
		FailOpenModeEnabled: false,
	}

	err = service.Update(testConfig)
	assert.Nil(t, err)

	// Verify the update
	updatedConfig, err := service.Get()
	assert.Nil(t, err)
	assert.NotNil(t, updatedConfig)
	assert.Equal(t, false, updatedConfig.Enabled)
	assert.Equal(t, url, *updatedConfig.URL)
	assert.Equal(t, username, *updatedConfig.Username)
	assert.Equal(t, timeout, *updatedConfig.TimeoutSeconds)

	// Restore original configuration
	// Note: Password is not returned by GET for security reasons, so we need to set it
	restoreConfig := iq.IQServerConfiguration{
		Enabled:             originalEnabled,
		ShowLink:            originalConfig.ShowLink,
		URL:                 originalURL,
		AuthenticationType:  originalAuthType,
		Username:            originalUsername,
		Password:            &password, // Use any password since we can't retrieve the original
		UseTrustStoreForURL: originalConfig.UseTrustStoreForURL,
		TimeoutSeconds:      originalConfig.TimeoutSeconds,
		Properties:          originalConfig.Properties,
		FailOpenModeEnabled: originalConfig.FailOpenModeEnabled,
	}

	// Only restore if there was a valid config originally
	if originalURL != nil && *originalURL != "" {
		err = service.Update(restoreConfig)
		assert.Nil(t, err)
	}
}

func TestIQServerEnableDisable(t *testing.T) {
	service := getTestService()

	// Get current configuration
	originalConfig, err := service.Get()
	assert.Nil(t, err)
	originalEnabled := originalConfig.Enabled

	// Set up a valid configuration first (with enabled=false)
	url := "https://iq-test.example.com"
	authType := "USER"
	username := "testuser"
	password := "testpass"
	timeout := 30

	testConfig := iq.IQServerConfiguration{
		Enabled:             false,
		ShowLink:            false,
		URL:                 &url,
		AuthenticationType:  &authType,
		Username:            &username,
		Password:            &password,
		UseTrustStoreForURL: false,
		TimeoutSeconds:      &timeout,
		Properties:          nil,
		FailOpenModeEnabled: false,
	}

	err = service.Update(testConfig)
	assert.Nil(t, err)

	// Test Disable (should already be disabled, but test the method)
	err = service.Disable()
	assert.Nil(t, err)

	config, err := service.Get()
	assert.Nil(t, err)
	assert.False(t, config.Enabled)

	// Note: We don't test Enable() because it would try to actually connect to the IQ Server
	// and fail if the URL is not reachable. In a real environment, Enable() would be tested
	// with a valid, reachable IQ Server.

	// Restore original state
	// Note: We need to provide the password again because GET doesn't return it
	originalConfig.Enabled = originalEnabled
	originalConfig.Password = &password
	err = service.Update(*originalConfig)
	assert.Nil(t, err)
}

func TestIQServerVerifyConnection(t *testing.T) {
	// Skip this test by default as it requires a valid IQ Server connection
	// In production, you would set up a valid IQ Server configuration first
	t.Skip("Skipping VerifyConnection test - requires valid IQ Server configuration")

	service := getTestService()

	// This would fail unless there's a valid IQ Server configured
	err := service.VerifyConnection()
	if err != nil {
		t.Logf("VerifyConnection failed (expected if no IQ Server is configured): %v", err)
	}
}
