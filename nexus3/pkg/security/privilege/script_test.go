package privilege_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/security/privilege"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/tools"
	"github.com/williamt1997/go-nexus-client/nexus3/schema"
	schemasecurity "github.com/williamt1997/go-nexus-client/nexus3/schema/security"
)

func getDefaultConfig() client.Config {
	return client.Config{
		Insecure: tools.GetEnv("NEXUS_INSECURE_SKIP_VERIFY", true).(bool),
		Password: tools.GetEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      tools.GetEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: tools.GetEnv("NEXUS_USRNAME", "admin").(string),
	}
}

var (
	testClient *client.Client = nil
)

func getTestPrivilegeScript(name string) *schema.Script {
	return &schema.Script{
		Name:    name,
		Content: fmt.Sprintf("log.info('Hello, %s!')", name),
		Type:    "groovy",
	}
}

func getTestClient() *client.Client {
	if testClient != nil {
		return testClient
	}
	return client.NewClient(getDefaultConfig())
}

func getTestService() *privilege.SecurityPrivilegeScriptService {
	return privilege.NewSecurityPrivilegeScriptService(getTestClient())
}

func getSecurityPrivilegeService() *privilege.SecurityPrivilegeService {
	return privilege.NewSecurityPrivilegeService(getTestClient())
}

func getTestScriptPrivilege(name string, description string, actions []schemasecurity.SecurityPrivilegeScriptActions, scriptName string) *schemasecurity.PrivilegeScript {
	return &schemasecurity.PrivilegeScript{
		Name:        name,
		Description: description,
		Actions:     actions,
		ScriptName:  scriptName,
	}
}

func TestScriptPrivilegeSecurity(t *testing.T) {
	privilegeScriptName := fmt.Sprintf("name-%d", tools.GetSeededRandomInteger(999))
	scriptName := fmt.Sprintf("script-%d", tools.GetSeededRandomInteger(999))
	testService := getTestService()
	securityPrivilegeService := getSecurityPrivilegeService()
	scriptService := nexus3.NewScriptService(getTestClient())
	scriptPrivilegePre := getTestScriptPrivilege(privilegeScriptName, "pre description", []schemasecurity.SecurityPrivilegeScriptActions{"BROWSE"}, scriptName)
	scriptPrivilegePost := getTestScriptPrivilege(privilegeScriptName, "post description", []schemasecurity.SecurityPrivilegeScriptActions{"BROWSE", "READ", "EDIT", "ADD", "DELETE", "RUN"}, scriptName)

	// Create a test script
	err := scriptService.Create(getTestPrivilegeScript(scriptName))
	assert.Nil(t, err)

	// Grant permissions to test script
	err = testService.Create(*scriptPrivilegePre)
	assert.Nil(t, err)

	// Get created privilege-script object and do some checks
	readPermission, err := securityPrivilegeService.Get(privilegeScriptName)
	assert.NoError(t, err)
	assert.Equal(t, privilegeScriptName, readPermission.Name)
	assert.Equal(t, "pre description", readPermission.Description)
	assert.Equal(t, []string{"BROWSE"}, readPermission.Actions)
	assert.Equal(t, scriptName, readPermission.ScriptName)

	// Update privilege-script object
	err = testService.Update(privilegeScriptName, *scriptPrivilegePost)
	assert.Nil(t, err)

	// Get updated object and test for updated values
	readPermission, err = securityPrivilegeService.Get(privilegeScriptName)
	assert.NoError(t, err)
	assert.Equal(t, privilegeScriptName, readPermission.Name)
	assert.Equal(t, "post description", readPermission.Description)
	assert.Equal(t, []string{"BROWSE", "READ", "EDIT", "ADD", "DELETE", "RUN"}, readPermission.Actions)
	assert.Equal(t, scriptName, readPermission.ScriptName)
}
