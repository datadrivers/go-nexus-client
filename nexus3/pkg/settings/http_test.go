package settings

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/settings"
	"github.com/stretchr/testify/assert"
)

func TestSettingsHTTP(t *testing.T) {
	service := getTestService()

	oldHTTPSettings, err := service.HTTP.Read()
	assert.Nil(t, err)
	assert.NotNil(t, oldHTTPSettings)

	timeout := int32(120)
	retries := int32(3)
	userAgent := "nexus-go-client-test/1.0"

	newHTTPSettings := settings.HTTPSettings{
		UserAgent: tools.GetStringPointer(userAgent),
		Timeout:   &timeout,
		Retries:   &retries,
		HTTPProxy: &settings.ProxySettings{
			Enabled: true,
			Host:    "proxy.internal",
			Port:    "3128",
			AuthInfo: &settings.AuthSettings{
				Enabled:  true,
				Username: "proxy-user",
				Password: "proxy-password",
			},
		},
		NonProxyHosts: []string{"localhost", "internal.registry.local"},
	}
	err = service.HTTP.Update(newHTTPSettings)
	assert.Nil(t, err)

	httpSettings, err := service.HTTP.Read()
	assert.Nil(t, err)
	assert.NotNil(t, httpSettings)
	assert.Equal(t, userAgent, *httpSettings.UserAgent)
	assert.Equal(t, timeout, *httpSettings.Timeout)
	assert.Equal(t, retries, *httpSettings.Retries)
	assert.NotNil(t, httpSettings.HTTPProxy)
	assert.True(t, httpSettings.HTTPProxy.Enabled)
	assert.Equal(t, "proxy.internal", httpSettings.HTTPProxy.Host)
	assert.Equal(t, "3128", httpSettings.HTTPProxy.Port)
	assert.ElementsMatch(t, newHTTPSettings.NonProxyHosts, httpSettings.NonProxyHosts)

	err = service.HTTP.Reset()
	assert.Nil(t, err)
}
