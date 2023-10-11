package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecurityRealmsActivate(t *testing.T) {
	service := getTestService()
	activeRealms, err := service.Realm.ListActive()
	assert.Nil(t, err)
	assert.NotNil(t, activeRealms)

	err = service.Realm.Activate(activeRealms)
	assert.Nil(t, err)
}

func TestSecurityRealmsActive(t *testing.T) {
	service := getTestService()
	activeRealms, err := service.Realm.ListActive()
	assert.Nil(t, err)
	assert.NotNil(t, activeRealms)
	assert.Greater(t, len(activeRealms), 0)
	assert.Contains(t, activeRealms, "NexusAuthenticatingRealm")
}

func TestSecurityRealmsAvailable(t *testing.T) {
	service := getTestService()
	availableRealms, err := service.Realm.ListAvailable()
	assert.Nil(t, err)
	assert.NotNil(t, availableRealms)
	assert.Greater(t, len(availableRealms), 0)

	for _, realm := range availableRealms {
		assert.NotNil(t, realm.ID)
		assert.NotNil(t, realm.Name)
	}
}
