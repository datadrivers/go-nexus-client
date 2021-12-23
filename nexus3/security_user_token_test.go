package nexus3

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func TestUserTokens(t *testing.T) {
	if getEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}
	client := getTestClient()

	userTokens := security.UserTokenConfiguration{
		Enabled:        true,
		ProtectContent: true,
	}
	err := client.Security.UserTokens.Configure(userTokens)
	assert.Nil(t, err)
	createdUserTokens, err := client.Security.UserTokens.Get()
	assert.Nil(t, err)
	assert.NotNil(t, createdUserTokens)
	assert.Equal(t, userTokens.Enabled, createdUserTokens.Enabled)
	assert.Equal(t, userTokens.ProtectContent, createdUserTokens.ProtectContent)

	createdUserTokens.Enabled = false
	createdUserTokens.ProtectContent = false
	err = client.Security.UserTokens.Configure(*createdUserTokens)
	assert.Nil(t, err)

	updatedUserTokens, err := client.Security.UserTokens.Get()
	assert.Nil(t, err)
	assert.NotNil(t, updatedUserTokens)
	assert.Equal(t, createdUserTokens.Enabled, updatedUserTokens.Enabled)
	assert.Equal(t, createdUserTokens.ProtectContent, updatedUserTokens.ProtectContent)

}
