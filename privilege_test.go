package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrivileges(t *testing.T) {
	client := getTestClient()

	privs, err := client.Privileges()
	assert.Nil(t, err)
	assert.NotNil(t, privs)
	assert.Greater(t, len(privs), 0)
}

func TestPrivilegeTypeWildcardRead(t *testing.T) {
	client := getTestClient()
	privName := "nx-all"

	priv, err := client.PrivilegeRead(privName)
	assert.Nil(t, err)
	assert.NotNil(t, priv)
	if priv != nil {
		assert.Equal(t, privName, priv.Name)
		assert.Equal(t, true, priv.ReadOnly)
		assert.Equal(t, "nexus:*", priv.Pattern)
		assert.Equal(t, "All permissions", priv.Description)
		assert.Equal(t, PrivilegeTypeWildcard, priv.Type)
		assert.Equal(t, 0, len(priv.Actions))
	}
}

func TestPrivilegeTypeScriptCreateReadAndDelete(t *testing.T) {
	client := getTestClient()
	testPrivilegeName := "test-script-privilege"
	testScriptName := "test-script"

	createScriptErr := client.ScriptCreate(&Script{
		Name:    testScriptName,
		Content: "log.info('Test a script privilege')",
		Type:    "groovy",
	})
	assert.Nil(t, createScriptErr)

	testScriptPrivilege := testScriptPrivilege(testPrivilegeName, testScriptName)
	createPrivilegeErr := client.PrivilegeCreate(testScriptPrivilege)
	assert.Nil(t, createPrivilegeErr)

	readPrivilege, readPrivilegeErr := client.PrivilegeRead(testPrivilegeName)
	assert.Nil(t, readPrivilegeErr)
	if readPrivilegeErr != nil {
		assert.Equal(t, testScriptPrivilege.Name, readPrivilege.Name)
		assert.Equal(t, testScriptPrivilege.Type, readPrivilege.Type)
		assert.Equal(t, testScriptPrivilege.ScriptName, readPrivilege.ScriptName)

		deletePrivilegeErr := client.PrivilegeDelete(testPrivilegeName)
		assert.Nil(t, deletePrivilegeErr)
		deleteScriptErr := client.ScriptDelete(testScriptName)
		assert.Nil(t, deleteScriptErr)
	}

	err := client.PrivilegeDelete(testPrivilegeName)
	assert.Nil(t, err)

	deletedPrivilege, err := client.PrivilegeRead(testScriptPrivilege.Name)
	assert.Nil(t, err)
	assert.Nil(t, deletedPrivilege)
	client.ScriptDelete(testScriptName)
}

func TestPrivilegeTypeAnalyticsRead(t *testing.T) {
	client := getTestClient()
	privName := "nx-analytics-all"

	priv, err := client.PrivilegeRead(privName)
	assert.Nil(t, err)
	assert.NotNil(t, priv)
	if priv != nil {
		assert.Equal(t, privName, priv.Name)
		assert.Equal(t, true, priv.ReadOnly)
		assert.Equal(t, "All permissions for Analytics", priv.Description)
		assert.Equal(t, PrivilegeTypeApplication, priv.Type)
		assert.Equal(t, 1, len(priv.Actions))
		assert.Equal(t, "ALL", priv.Actions[0])
		// Attributes of other types
		assert.Equal(t, "", priv.Format)
		assert.Equal(t, "", priv.Repository)
	}
}

func TestPrivilegeTypeApplicationRead(t *testing.T) {
	client := getTestClient()
	privName := "nx-apikey-all"

	priv, err := client.PrivilegeRead(privName)
	assert.Nil(t, err)
	assert.NotNil(t, priv)
	if priv != nil {
		assert.Equal(t, privName, priv.Name)
		assert.Equal(t, true, priv.ReadOnly)
		assert.Equal(t, "All permissions for APIKey", priv.Description)
		assert.Equal(t, PrivilegeTypeApplication, priv.Type)
		assert.Equal(t, 1, len(priv.Actions))
		assert.Equal(t, "ALL", priv.Actions[0])
		// Attributes of other types
		assert.Equal(t, "", priv.Format)
		assert.Equal(t, "", priv.Repository)
	}
}

func TestPrivilegeTypeRepositoryAdminRead(t *testing.T) {
	client := getTestClient()
	privName := "nx-repository-admin-*-*-*"

	priv, err := client.PrivilegeRead(privName)
	assert.Nil(t, err)
	assert.NotNil(t, priv)
	if priv != nil {
		assert.Equal(t, privName, priv.Name)
		assert.Equal(t, true, priv.ReadOnly)
		assert.Equal(t, "All privileges for all repository administration", priv.Description)
		assert.Equal(t, PrivilegeTypeRepositoryAdmin, priv.Type)
		assert.Equal(t, 1, len(priv.Actions))
		assert.Equal(t, "ALL", priv.Actions[0])
		assert.Equal(t, "*", priv.Format)
		assert.Equal(t, "*", priv.Repository)
	}
}

func TestPrivilegeTypeRepositoryViewRead(t *testing.T) {
	client := getTestClient()
	privName := "nx-repository-view-*-*-*"

	priv, err := client.PrivilegeRead(privName)
	assert.Nil(t, err)
	assert.NotNil(t, priv)
	if priv != nil {
		assert.Equal(t, privName, priv.Name)
		assert.Equal(t, true, priv.ReadOnly)
		assert.Equal(t, "All permissions for all repository views", priv.Description)
		assert.Equal(t, PrivilegeTypeRepositoryView, priv.Type)
		assert.Equal(t, 1, len(priv.Actions))
		assert.Equal(t, "ALL", priv.Actions[0])
		assert.Equal(t, "*", priv.Format)
		assert.Equal(t, "*", priv.Repository)
	}
}

func TestPrivilegeCreateReadUpdateDelete(t *testing.T) {
	client := getTestClient()
	privilege := testPrivilege("test-privilege")

	err := client.PrivilegeCreate(privilege)
	assert.Nil(t, err)

	createdPrivilege, err := client.PrivilegeRead(privilege.Name)
	assert.Nil(t, err)
	assert.NotNil(t, createdPrivilege)

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
	assert.NotNil(t, updatedPrivilege)
	assert.Equal(t, createdPrivilege.Description, updatedPrivilege.Description)
	assert.Equal(t, createdPrivilege.Domain, updatedPrivilege.Domain)

	err = client.PrivilegeDelete(privilege.Name)
	assert.Nil(t, err)

	deletedPrivilege, err := client.PrivilegeRead(privilege.Name)
	assert.Nil(t, err)
	assert.Nil(t, deletedPrivilege)
}

func testScriptPrivilege(name string, scriptName string) Privilege {
	return Privilege{
		Actions:     []string{"READ"},
		Name:        name,
		Description: "Description",
		ScriptName:  scriptName,
		Type:        PrivilegeTypeScript,
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
