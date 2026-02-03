package privilege_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/security/privilege"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/tools"
	schemasecurity "github.com/williamt1997/go-nexus-client/nexus3/schema/security"
)

func getTestPrivilegeApplication(name string, description string, actions []schemasecurity.SecurityPrivilegeApplicationActions, domain string) *schemasecurity.PrivilegeApplication {
	return &schemasecurity.PrivilegeApplication{
		Name:        name,
		Description: description,
		Actions:     actions,
		Domain:      domain,
	}
}

func TestApplicationPrivilegeSecurity(t *testing.T) {
	PrivilegeApplicationName := fmt.Sprintf("application-%d", tools.GetSeededRandomInteger(999))
	testService := privilege.NewSecurityPrivilegeApplicationService(getTestClient())
	privilegeService := privilege.NewSecurityPrivilegeService(getTestClient())

	// Create application-privilege object
	applicationPrivilege := getTestPrivilegeApplication(PrivilegeApplicationName, "demo descrp", []schemasecurity.SecurityPrivilegeApplicationActions{"READ"}, "domain")
	err := testService.Create(*applicationPrivilege)
	assert.Nil(t, err)

	// Fetch recently created application-privilege object and do some checks
	applicationPrivilegeFetched, err := privilegeService.Get(PrivilegeApplicationName)
	assert.Nil(t, err)
	assert.Equal(t, PrivilegeApplicationName, applicationPrivilegeFetched.Name)
	assert.Equal(t, "demo descrp", applicationPrivilegeFetched.Description)
	assert.Equal(t, []string{"READ"}, applicationPrivilegeFetched.Actions)

	// Update application-privilege object
	applicationPrivilege = getTestPrivilegeApplication(PrivilegeApplicationName, "demo descrp updated", []schemasecurity.SecurityPrivilegeApplicationActions{"ADD", "READ", "DELETE", "ASSOCIATE"}, "domain")
	err = testService.Update(PrivilegeApplicationName, *applicationPrivilege)
	assert.Nil(t, err)
	applicationPrivilegeFetched, err = privilegeService.Get(PrivilegeApplicationName)
	assert.Nil(t, err)
	assert.Equal(t, "demo descrp updated", applicationPrivilegeFetched.Description)
	assert.Equal(t, []string{"ADD", "READ", "DELETE", "ASSOCIATE"}, applicationPrivilegeFetched.Actions)
	assert.Equal(t, "domain", applicationPrivilege.Domain)

	// Delete application-privilege-object
	err = privilegeService.Delete(PrivilegeApplicationName)
	assert.Nil(t, err)

	// Check for successful deletion
	applicationPrivilegeFetched, err = privilegeService.Get(PrivilegeApplicationName)
	assert.Error(t, err)
	assert.Nil(t, applicationPrivilegeFetched)
}
