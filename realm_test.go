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
	assert.Greater(t, len(activeRealms), 0)
	assert.Contains(t, activeRealms, "NexusAuthenticatingRealm")
	assert.Contains(t, activeRealms, "NexusAuthorizingRealm")
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
