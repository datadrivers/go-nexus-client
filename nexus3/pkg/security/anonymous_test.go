package security

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func TestSecurityAnonymous(t *testing.T) {
	service := getTestService()

	oldAnonymous, err := service.Anonymous.Read()
	assert.Nil(t, err)
	assert.NotNil(t, oldAnonymous)

	newAnonymous := security.AnonymousAccessSettings{
		Enabled:   true,
		UserID:    "anonymous",
		RealmName: "NexusAuthorizingRealm",
	}
	err = service.Anonymous.Update(newAnonymous)
	assert.Nil(t, err)

	anonymous, err := service.Anonymous.Read()
	assert.Nil(t, err)
	assert.NotNil(t, anonymous)
	assert.Equal(t, *anonymous, newAnonymous)

	err = service.Anonymous.Update(*oldAnonymous)
	assert.Nil(t, err)
}
