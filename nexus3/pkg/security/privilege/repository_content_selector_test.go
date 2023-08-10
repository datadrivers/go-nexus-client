package privilege_test

import (
	"fmt"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/security"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/security/privilege"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	schemasecurity "github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func getTestPrivilegeRepositoryContentSelector(name string, description string, actions []string, format string, repository string, contentSelector string) *schemasecurity.PrivilegeRepositoryContentSelector {
	return &schemasecurity.PrivilegeRepositoryContentSelector{
		Name:            name,
		Description:     description,
		Actions:         actions,
		Format:          format,
		Repository:      repository,
		ContentSelector: contentSelector,
	}
}

func getTestContentSelector(name string, description string, expression string) *schemasecurity.ContentSelector {
	return &schemasecurity.ContentSelector{
		Name:        name,
		Description: description,
		Expression:  expression,
	}
}

func TestContentSelectorPrivilegeSecurity(t *testing.T) {
	contentSelectorName := fmt.Sprintf("content-selector-%d", tools.GetSeededRandomInteger(999))
	contentSelectorExpression := `format == "npm" or (format == "maven2" and path =~ "^/org/apache/commons/.*")`
	privilegeRepositoryContentSelectorName := fmt.Sprintf("content-selector-privilege%d", tools.GetSeededRandomInteger(999))
	testService := privilege.NewSecurityPrivilegeContentSelectorService(getTestClient())
	contentSelectorService := security.NewSecurityContentSelectorService(getTestClient())
	privilegeService := getSecurityPrivilegeService()

	// Create Content Selector Object
	err := contentSelectorService.Create(*getTestContentSelector(contentSelectorName, "description", contentSelectorExpression))
	assert.Nil(t, err)

	// Create repository-content-selector-privilege object for already existing Maven repo (was created by Nexus itself)
	contentSelectorPrivilege := getTestPrivilegeRepositoryContentSelector(privilegeRepositoryContentSelectorName, "descr", []string{"ADD"}, "maven2", "maven-snapshots", contentSelectorName)
	err = testService.Create(*contentSelectorPrivilege)
	assert.Nil(t, err)

	// Fetch recently created repository-content-selector-privilege object and do some checks
	contentSelectorPrivilegeFetched, err := privilegeService.Get(privilegeRepositoryContentSelectorName)
	assert.Nil(t, err)
	assert.Equal(t, privilegeRepositoryContentSelectorName, contentSelectorPrivilegeFetched.Name)
	assert.Equal(t, "descr", contentSelectorPrivilegeFetched.Description)
	assert.Equal(t, []string{"ADD"}, contentSelectorPrivilegeFetched.Actions)
	assert.Equal(t, "maven2", contentSelectorPrivilegeFetched.Format)
	assert.Equal(t, "maven-snapshots", contentSelectorPrivilegeFetched.Repository)
	assert.Equal(t, contentSelectorName, contentSelectorPrivilegeFetched.ContentSelector)

	// Update repository-content-selector-privilege object
	contentSelectorPrivilege = getTestPrivilegeRepositoryContentSelector(privilegeRepositoryContentSelectorName, "demo descrp", []string{"BROWSE", "READ", "EDIT", "ADD", "DELETE"}, "maven2", "maven-snapshots", contentSelectorName)
	err = testService.Update(privilegeRepositoryContentSelectorName, *contentSelectorPrivilege)
	assert.Nil(t, err)
	contentSelectorPrivilegeFetched, err = privilegeService.Get(privilegeRepositoryContentSelectorName)
	assert.Nil(t, err)
	assert.Equal(t, privilegeRepositoryContentSelectorName, contentSelectorPrivilegeFetched.Name)
	assert.Equal(t, []string{"BROWSE", "READ", "EDIT", "ADD", "DELETE"}, contentSelectorPrivilegeFetched.Actions)

	// // Delete repository-content-selector-privilege-object
	err = privilegeService.Delete(privilegeRepositoryContentSelectorName)
	assert.Nil(t, err)

	// Check for successful deletion
	contentSelectorPrivilegeFetched, err = privilegeService.Get(privilegeRepositoryContentSelectorName)
	assert.Error(t, err)
	assert.Nil(t, contentSelectorPrivilegeFetched)
}
