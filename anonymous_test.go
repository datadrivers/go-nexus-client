package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymous(t *testing.T) {
	client := getTestClient()

	oldAnonymous, err := client.AnonymousRead()
	assert.Nil(t, err)
	assert.NotNil(t, oldAnonymous)

	newAnonymous := AnonymousConfig{
		Enabled:   true,
		UserID:    "anonymous",
		RealmName: "NexusAuthorizingRealm",
	}
	err = client.AnonymousUpdate(newAnonymous)
	assert.Nil(t, err)

	anonymous, err := client.AnonymousRead()
	assert.Nil(t, err)
	assert.NotNil(t, anonymous)
	assert.Equal(t, *anonymous, newAnonymous)

	err = client.AnonymousUpdate(*oldAnonymous)
	assert.Nil(t, err)
}
