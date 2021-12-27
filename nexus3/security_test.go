package nexus3

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func TestPrivilegeTypeScriptCreateReadAndDelete(t *testing.T) {
	client := getTestClient()
	testPrivilegeName := "test-script-privilege"
	testScriptName := "test-script"

	createScriptErr := client.Script.Create(&schema.Script{
		Name:    testScriptName,
		Content: "log.info('Test a script privilege')",
		Type:    "groovy",
	})
	assert.Nil(t, createScriptErr)

	testScriptPrivilege := testScriptPrivilege(testPrivilegeName, testScriptName)
	createPrivilegeErr := client.Security.Privilege.Create(testScriptPrivilege)
	assert.Nil(t, createPrivilegeErr)

	readPrivilege, readPrivilegeErr := client.Security.Privilege.Get(testScriptPrivilege.Name)
	assert.Nil(t, readPrivilegeErr)
	assert.Equal(t, testScriptPrivilege.Name, readPrivilege.Name)
	assert.Equal(t, testScriptPrivilege.Type, readPrivilege.Type)
	assert.Equal(t, testScriptPrivilege.ScriptName, readPrivilege.ScriptName)

	deletePrivilegeErr := client.Security.Privilege.Delete(testScriptPrivilege.Name)
	assert.Nil(t, deletePrivilegeErr)
	deleteScriptErr := client.Script.Delete(testScriptName)
	assert.Nil(t, deleteScriptErr)

	deletedPrivilege, err := client.Security.Privilege.Get(testScriptPrivilege.Name)
	assert.Nil(t, err)
	assert.Nil(t, deletedPrivilege)
	client.Script.Delete(testScriptName)
}

func testScriptPrivilege(name string, scriptName string) security.Privilege {
	return security.Privilege{
		Actions:     []string{"READ"},
		Name:        name,
		Description: "Description",
		ScriptName:  scriptName,
		Type:        security.PrivilegeTypeScript,
	}
}
