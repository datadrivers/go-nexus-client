package security

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func TestUserTokens(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}
	service := getTestService()

	userTokens := security.UserTokenConfiguration{
		Enabled:        true,
		ProtectContent: true,
		ExpirationDays: int(45),
	}
	err := service.UserTokens.Configure(userTokens)
	assert.Nil(t, err)
	createdUserTokens, err := service.UserTokens.Get()
	assert.Nil(t, err)
	assert.NotNil(t, createdUserTokens)
	assert.Equal(t, userTokens.ExpirationDays, createdUserTokens.ExpirationDays)
	assert.Equal(t, userTokens.Enabled, createdUserTokens.Enabled)
	assert.Equal(t, userTokens.ProtectContent, createdUserTokens.ProtectContent)

	createdUserTokens.ExpirationDays = int(30)
	createdUserTokens.ProtectContent = false
	err = service.UserTokens.Configure(*createdUserTokens)
	assert.Nil(t, err)

	updatedUserTokens, err := service.UserTokens.Get()
	assert.Nil(t, err)
	assert.NotNil(t, updatedUserTokens)
	assert.Equal(t, createdUserTokens.ExpirationDays, updatedUserTokens.ExpirationDays)
	assert.Equal(t, createdUserTokens.ProtectContent, updatedUserTokens.ProtectContent)

}
