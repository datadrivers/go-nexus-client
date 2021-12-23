package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserTokens(t *testing.T) {
	if getEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}
	client := getTestClient()

	userTokens := UserTokenConfiguration{
		Enabled:        true,
		ProtectContent: true,
	}
	err := client.UserTokensApply(userTokens)
	assert.Nil(t, err)

	if err == nil {
		createdUserTokens, err := client.UserTokensRead()
		assert.Nil(t, err)
		assert.NotNil(t, createdUserTokens)
		assert.Equal(t, userTokens.Enabled, createdUserTokens.Enabled)
		assert.Equal(t, userTokens.ProtectContent, createdUserTokens.ProtectContent)

		createdUserTokens.Enabled = false
		createdUserTokens.ProtectContent = false
		err = client.UserTokensApply(*createdUserTokens)
		assert.Nil(t, err)

		updatedUserTokens, err := client.UserTokensRead()
		assert.Nil(t, err)
		assert.NotNil(t, updatedUserTokens)
		assert.Equal(t, createdUserTokens.Enabled, updatedUserTokens.Enabled)
		assert.Equal(t, createdUserTokens.ProtectContent, updatedUserTokens.ProtectContent)
	}
}
