package privilege_test

import (
	"fmt"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/security/privilege"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	schemasecurity "github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func getTestPrivilegeRepository(name string, description string, actions []string, format string, repository string) *schemasecurity.PrivilegeRepositoryView {
	return &schemasecurity.PrivilegeRepositoryView{
		Name:        name,
		Description: description,
		Actions:     actions,
		Format:      format,
		Repository:  repository,
	}
}

func TestRepositoryPrivilegeSecurity(t *testing.T) {
	privilegeRepositoryName := fmt.Sprintf("repository-%d", tools.GetSeededRandomInteger(999))
	testService := privilege.NewSecurityPrivilegeRepositoryViewService(getTestClient())
	privilegeService := privilege.NewSecurityPrivilegeService(getTestClient())

	// Create repository-privilege object for already existing Maven repo (was created by Nexus itself)
	repositoryPrivilege := getTestPrivilegeRepository(privilegeRepositoryName, "demo descrp", []string{"BROWSE", "READ"}, "maven2", "maven-snapshots")

	err := testService.Create(*repositoryPrivilege)
	assert.Nil(t, err)

	// Fetch recently created repository-privilege object and do some checks
	repositoryPrivilegeFetched, err := privilegeService.Get(privilegeRepositoryName)
	assert.Nil(t, err)
	assert.Equal(t, privilegeRepositoryName, repositoryPrivilegeFetched.Name)
	assert.Equal(t, "demo descrp", repositoryPrivilegeFetched.Description)
	assert.Equal(t, []string{"BROWSE", "READ"}, repositoryPrivilegeFetched.Actions)
	assert.Equal(t, "maven2", repositoryPrivilegeFetched.Format)
	assert.Equal(t, "maven-snapshots", repositoryPrivilegeFetched.Repository)

	// Update repository-privilege object
	repositoryPrivilege = getTestPrivilegeRepository(privilegeRepositoryName, "demo descrp", []string{"BROWSE", "READ", "EDIT", "ADD", "DELETE"}, "maven2", "maven-snapshots")
	err = testService.Update(privilegeRepositoryName, *repositoryPrivilege)
	assert.Nil(t, err)
	repositoryPrivilegeFetched, err = privilegeService.Get(privilegeRepositoryName)
	assert.Nil(t, err)
	assert.Equal(t, privilegeRepositoryName, repositoryPrivilegeFetched.Name)
	assert.Equal(t, []string{"BROWSE", "READ", "EDIT", "ADD", "DELETE"}, repositoryPrivilegeFetched.Actions)

	// Delete repository-privilege-object
	err = privilegeService.Delete(privilegeRepositoryName)
	assert.Nil(t, err)

	// Check for successful deletion
	repositoryPrivilegeFetched, err = privilegeService.Get(privilegeRepositoryName)
	assert.Error(t, err)
	assert.Nil(t, repositoryPrivilegeFetched)
}
