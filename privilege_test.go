package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrivilegeCreateReadUpdateDelete(t *testing.T) {
	client := NewClient(getDefaultConfig())
	privilege := testPrivilege("test-privilege")

	err := client.PrivilegeCreate(privilege)
	assert.Nil(t, err)

	if err != nil {
		createdPrivilege, err := client.PrivilegeRead(privilege.Name)
		assert.Nil(t, err)
		assert.Equal(t, privilege.Name, createdPrivilege.Name)
		assert.Equal(t, privilege.Description, createdPrivilege.Description)
		assert.Equal(t, privilege.Domain, createdPrivilege.Domain)
		assert.Equal(t, privilege.Type, createdPrivilege.Type)

		// Update
		createdPrivilege.Description = "updated"
		createdPrivilege.Domain = "datastores"

		err = client.PrivilegeUpdate(privilege.Name, *createdPrivilege)
		assert.Nil(t, err)
		updatedPrivilege, err := client.PrivilegeRead(privilege.Name)
		assert.Nil(t, err)
		assert.Equal(t, createdPrivilege.Description, updatedPrivilege.Description)
		assert.Equal(t, createdPrivilege.Domain, updatedPrivilege.Domain)

		err = client.PrivilegeDelete(privilege.Name)
		assert.Nil(t, err)

		deletedPrivilege, err := client.PrivilegeRead(privilege.Name)
		assert.Nil(t, err)
		assert.Nil(t, deletedPrivilege)
	}
}

func testPrivilege(name string) Privilege {
	return Privilege{
		Actions:     []string{"READ"},
		Description: fmt.Sprintf("Go client privilege %d", time.Now().Unix()),
		Domain:      "*",
		Name:        name,
		Type:        "application",
	}
}
