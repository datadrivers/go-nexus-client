package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleCreate(t *testing.T) {
	client := NewClient(getDefaultConfig())
	role := testRole("test-role-create")

	err := client.RoleCreate(*role)
	assert.Nil(t, err)

	err = client.RoleDelete(role.ID)
	assert.Nil(t, err)
}

func TestRoleRead(t *testing.T) {
	client := NewClient(getDefaultConfig())
	role, err := client.RoleRead("nx-admin")

	assert.Nil(t, err)
	assert.NotNil(t, role)
	assert.Equal(t, role.ID, "nx-admin")
	assert.Equal(t, role.Name, "nx-admin")
}

func TestRoleCreateReadUpdateDelete(t *testing.T) {
	client := NewClient(getDefaultConfig())
	testRole := testRole("test-role-create-read-update-delete")

	// Create
	err := client.RoleCreate(*testRole)
	assert.Nil(t, err)

	// Read
	role, err := client.RoleRead(testRole.ID)
	assert.Nil(t, err)
	assert.NotNil(t, role)
	assert.Equal(t, testRole.ID, role.ID)
	assert.Equal(t, testRole.Name, role.Name)
	assert.Equal(t, testRole.Description, role.Description)

	// Update
	updatedRole := role
	updatedRole.Description = "changed"
	updatedRole.Name = "changed"
	updatedRole.Privileges = []string{"nx-repository-view-*-*-*"}
	updatedRole.Roles = []string{"nx-anonymous"}

	err = client.RoleUpdate(role.ID, *updatedRole)
	assert.Nil(t, err)

	updatedRole, err = client.RoleRead(updatedRole.ID)
	assert.Nil(t, err)
	assert.NotNil(t, updatedRole)
	assert.Equal(t, "changed", updatedRole.Description)
	assert.Equal(t, "changed", updatedRole.Name)
	assert.Equal(t, []string{"nx-repository-view-*-*-*"}, updatedRole.Privileges)
	assert.Equal(t, []string{"nx-anonymous"}, updatedRole.Roles)

	// Delete
	err = client.RoleDelete(role.ID)
	assert.Nil(t, err)

	role, err = client.RoleRead(testRole.ID)
	assert.Nil(t, err)
	assert.Nil(t, role)
}

func testRole(id string) *Role {
	return &Role{
		ID:          id,
		Name:        "nx-test-role",
		Description: "Go client test role",
		Privileges:  []string{"nx-all"},
		Roles:       []string{"nx-admin"},
	}
}
