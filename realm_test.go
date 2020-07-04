package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRealmsActivate(t *testing.T) {
	client := getTestClient()
	activeRealms, err := client.RealmsActive()
	assert.Nil(t, err)
	assert.NotNil(t, activeRealms)

	err = client.RealmsActivate(activeRealms)
	assert.Nil(t, err)
}

func TestRealmsActive(t *testing.T) {
	client := getTestClient()
	activeRealms, err := client.RealmsActive()
	assert.Nil(t, err)
	assert.NotNil(t, activeRealms)
	assert.Equal(t, 2, len(activeRealms))
	assert.Equal(t, "NexusAuthenticatingRealm", activeRealms[0])
	assert.Equal(t, "NexusAuthorizingRealm", activeRealms[1])
}

func TestRealmsAvailable(t *testing.T) {
	client := getTestClient()
	availableRealms, err := client.RealmsAvailable()
	assert.Nil(t, err)
	assert.NotNil(t, availableRealms)
	assert.Greater(t, len(availableRealms), 0)

	for _, realm := range availableRealms {
		assert.NotNil(t, realm.ID)
		assert.NotNil(t, realm.Name)
	}
}
