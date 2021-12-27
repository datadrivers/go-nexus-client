package security

import (
	"fmt"
	"testing"
	"time"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func TestPrivileges(t *testing.T) {
	service := getTestService()

	privs, err := service.Privilege.List()
	assert.Nil(t, err)
	assert.NotNil(t, privs)
	assert.Greater(t, len(privs), 0)
}

func TestPrivilegeTypeWildcardRead(t *testing.T) {
	service := getTestService()
	privName := "nx-all"

	priv, err := service.Privilege.Get(privName)
	assert.Nil(t, err)
	assert.NotNil(t, priv)
	if priv != nil {
		assert.Equal(t, privName, priv.Name)
		assert.Equal(t, true, priv.ReadOnly)
		assert.Equal(t, "nexus:*", priv.Pattern)
		assert.Equal(t, "All permissions", priv.Description)
		assert.Equal(t, security.PrivilegeTypeWildcard, priv.Type)
		assert.Equal(t, 0, len(priv.Actions))
	}
}

func TestPrivilegeTypeAnalyticsRead(t *testing.T) {
	service := getTestService()
	privName := "nx-analytics-all"

	priv, err := service.Privilege.Get(privName)
	assert.Nil(t, err)
	assert.NotNil(t, priv)
	if priv != nil {
		assert.Equal(t, privName, priv.Name)
		assert.Equal(t, true, priv.ReadOnly)
		assert.Equal(t, "All permissions for Analytics", priv.Description)
		assert.Equal(t, security.PrivilegeTypeApplication, priv.Type)
		assert.Equal(t, 1, len(priv.Actions))
		assert.Equal(t, "ALL", priv.Actions[0])
		// Attributes of other types
		assert.Equal(t, "", priv.Format)
		assert.Equal(t, "", priv.Repository)
	}
}

func TestPrivilegeTypeApplicationRead(t *testing.T) {
	service := getTestService()
	privName := "nx-apikey-all"

	priv, err := service.Privilege.Get(privName)
	assert.Nil(t, err)
	assert.NotNil(t, priv)
	if priv != nil {
		assert.Equal(t, privName, priv.Name)
		assert.Equal(t, true, priv.ReadOnly)
		assert.Equal(t, "All permissions for APIKey", priv.Description)
		assert.Equal(t, security.PrivilegeTypeApplication, priv.Type)
		assert.Equal(t, 1, len(priv.Actions))
		assert.Equal(t, "ALL", priv.Actions[0])
		// Attributes of other types
		assert.Equal(t, "", priv.Format)
		assert.Equal(t, "", priv.Repository)
	}
}

func TestPrivilegeTypeRepositoryAdminRead(t *testing.T) {
	service := getTestService()
	privName := "nx-repository-admin-*-*-*"

	priv, err := service.Privilege.Get(privName)
	assert.Nil(t, err)
	assert.NotNil(t, priv)
	if priv != nil {
		assert.Equal(t, privName, priv.Name)
		assert.Equal(t, true, priv.ReadOnly)
		assert.Equal(t, "All privileges for all repository administration", priv.Description)
		assert.Equal(t, security.PrivilegeTypeRepositoryAdmin, priv.Type)
		assert.Equal(t, 1, len(priv.Actions))
		assert.Equal(t, "ALL", priv.Actions[0])
		assert.Equal(t, "*", priv.Format)
		assert.Equal(t, "*", priv.Repository)
	}
}

func TestPrivilegeTypeRepositoryViewRead(t *testing.T) {
	service := getTestService()
	privName := "nx-repository-view-*-*-*"

	priv, err := service.Privilege.Get(privName)
	assert.Nil(t, err)
	assert.NotNil(t, priv)
	if priv != nil {
		assert.Equal(t, privName, priv.Name)
		assert.Equal(t, true, priv.ReadOnly)
		assert.Equal(t, "All permissions for all repository views", priv.Description)
		assert.Equal(t, security.PrivilegeTypeRepositoryView, priv.Type)
		assert.Equal(t, 1, len(priv.Actions))
		assert.Equal(t, "ALL", priv.Actions[0])
		assert.Equal(t, "*", priv.Format)
		assert.Equal(t, "*", priv.Repository)
	}
}

func TestPrivilegeCreateReadUpdateDelete(t *testing.T) {
	service := getTestService()
	privilege := testPrivilege("test-privilege")

	err := service.Privilege.Create(privilege)
	assert.Nil(t, err)

	createdPrivilege, err := service.Privilege.Get(privilege.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdPrivilege)

	assert.Equal(t, privilege.Name, createdPrivilege.Name)
	assert.Equal(t, privilege.Description, createdPrivilege.Description)
	assert.Equal(t, privilege.Domain, createdPrivilege.Domain)
	assert.Equal(t, privilege.Type, createdPrivilege.Type)

	// Update
	createdPrivilege.Description = "updated"
	createdPrivilege.Domain = "datastores"

	err = service.Privilege.Update(privilege.Name, *createdPrivilege)
	assert.Nil(t, err)

	updatedPrivilege, err := service.Privilege.Get(privilege.Name)
	assert.Nil(t, err)
	assert.NotNil(t, updatedPrivilege)
	assert.Equal(t, createdPrivilege.Description, updatedPrivilege.Description)
	assert.Equal(t, createdPrivilege.Domain, updatedPrivilege.Domain)

	err = service.Privilege.Delete(privilege.Name)
	assert.Nil(t, err)

	deletedPrivilege, err := service.Privilege.Get(privilege.Name)
	assert.Nil(t, err)
	assert.Nil(t, deletedPrivilege)
}

func testPrivilege(name string) security.Privilege {
	return security.Privilege{
		Actions:     []string{"READ"},
		Description: fmt.Sprintf("Go client privilege %d", time.Now().Unix()),
		Domain:      "*",
		Name:        name,
		Type:        "application",
	}
}

// Testfunction TestPrivilegeTypeScriptCreateReadAndDelete located in main client folder, because this test function use multiple services
