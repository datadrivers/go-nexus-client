package privilege_test

import (
	"fmt"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/security/privilege"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	schemasecurity "github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func getTestPrivilegeWildcard(name string, description string, pattern string) *schemasecurity.PrivilegeWildcard {
	return &schemasecurity.PrivilegeWildcard{
		Name:        name,
		Description: description,
		Pattern:     pattern,
	}
}

func TestWildcardPrivilegeSecurity(t *testing.T) {
	privilegeWildcardName := fmt.Sprintf("wildcard-%d", tools.GetSeededRandomInteger(999))
	testService := privilege.NewSecurityPrivilegeWildcardService(getTestClient())
	privilegeService := privilege.NewSecurityPrivilegeService(getTestClient())

	// Create wildcard-privilege object
	wildcardPrivilege := getTestPrivilegeWildcard(privilegeWildcardName, "demo descrp", "nexus:*")
	err := testService.Create(*wildcardPrivilege)
	assert.Nil(t, err)

	// Fetch recently created wildcard-privilege object and do some checks
	wildcardPrivilegeFetched, err := privilegeService.Get(privilegeWildcardName)
	assert.Nil(t, err)
	assert.Equal(t, privilegeWildcardName, wildcardPrivilegeFetched.Name)
	assert.Equal(t, "demo descrp", wildcardPrivilegeFetched.Description)
	assert.Equal(t, "nexus:*", wildcardPrivilegeFetched.Pattern)

	// Update wildcard-privilege object
	wildcardPrivilege = getTestPrivilegeWildcard(privilegeWildcardName, "demo descrp updated", "nexus:nexus")
	err = testService.Update(privilegeWildcardName, *wildcardPrivilege)
	assert.Nil(t, err)
	wildcardPrivilegeFetched, err = privilegeService.Get(privilegeWildcardName)
	assert.Nil(t, err)
	assert.Equal(t, "demo descrp updated", wildcardPrivilegeFetched.Description)
	assert.Equal(t, "nexus:nexus", wildcardPrivilegeFetched.Pattern)

	// Delete wildcard-privilege-object
	err = privilegeService.Delete(privilegeWildcardName)
	assert.Nil(t, err)

	// Check for successful deletion
	wildcardPrivilegeFetched, err = privilegeService.Get(privilegeWildcardName)
	assert.Error(t, err)
	assert.Nil(t, wildcardPrivilegeFetched)
}
