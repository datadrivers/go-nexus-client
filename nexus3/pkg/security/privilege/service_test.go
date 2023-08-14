package privilege

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

var (
	testClient *client.Client = nil
)

func getTestClient() *client.Client {
	if testClient != nil {
		return testClient
	}
	return client.NewClient(getDefaultConfig())
}

func getDefaultConfig() client.Config {
	return client.Config{
		Insecure: tools.GetEnv("NEXUS_INSECURE_SKIP_VERIFY", true).(bool),
		Password: tools.GetEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      tools.GetEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: tools.GetEnv("NEXUS_USRNAME", "admin").(string),
	}
}

func getTestService() *SecurityPrivilegeService {
	return NewSecurityPrivilegeService(getTestClient())
}

func TestNewPrivilegeService(t *testing.T) {
	s := getTestService()

	assert.NotNil(t, s, "NewPrivilegeService() must not return nil")
}

func TestListPrivileges(t *testing.T) {
	service := getTestService()
	privileges, err := service.List()

	assert.Nil(t, err)
	assert.NotEmpty(t, privileges)
}
func TestGetPrivilegesForDomain(t *testing.T) {
	service := getTestService()
	privilege, err := service.Get("nx-wonderland-all")

	assert.Nil(t, err)
	assert.NotNil(t, privilege)
	assert.Equal(t, "nx-wonderland-all", privilege.Name)
	assert.Equal(t, "All permissions for Wonderland", privilege.Description)
	assert.Equal(t, []string{"ALL"}, privilege.Actions)
	assert.Equal(t, "wonderland", privilege.Domain)
}

func TestGetPrivilegesForScript(t *testing.T) {
	service := getTestService()
	privilege, err := service.Get("nx-script-*-add")

	assert.Nil(t, err)
	assert.NotNil(t, privilege)
	assert.Equal(t, "nx-script-*-add", privilege.Name)
	assert.Equal(t, "Add permissions for Scripts", privilege.Description)
	assert.Equal(t, []string{"ADD", "READ"}, privilege.Actions)
	assert.NotEqual(t, []string{"READ", "ADD"}, privilege.Actions)
	assert.NotEqual(t, []string{"READ"}, privilege.Actions)
	assert.Equal(t, "*", privilege.ScriptName)
}

func TestGetPrivilegesForRepository(t *testing.T) {
	service := getTestService()
	privilege, err := service.Get("nx-repository-view-yum-*-browse")

	assert.Nil(t, err)
	assert.NotNil(t, privilege)
	assert.Equal(t, "nx-repository-view-yum-*-browse", privilege.Name)
	assert.Equal(t, "Browse privilege for all 'yum'-format repository views", privilege.Description)
	assert.Equal(t, []string{"BROWSE"}, privilege.Actions)
	assert.Equal(t, "*", privilege.Repository)
}

func TestPrivilegeActionIsValid(t *testing.T) {
	const errorMessage = "Invalid action privilege. See https://help.sonatype.com/repomanager3/nexus-repository-administration/access-control/privileges#Privileges-PrivilegeActions for possible values"
	var action security.ActionPrivilege

	action = "add"
	assert.Error(t, action.IsValid(), errorMessage)
	action = "add   "
	assert.Error(t, action.IsValid(), errorMessage)
	action = ""
	assert.Error(t, action.IsValid(), errorMessage)

	action = "ADD"
	assert.Nil(t, action.IsValid(), errorMessage)
	action = "BROWSE"
	assert.Nil(t, action.IsValid(), errorMessage)
	action = "CREATE"
	assert.Nil(t, action.IsValid(), errorMessage)
	action = "DELETE"
	assert.Nil(t, action.IsValid(), errorMessage)
	action = "EDIT"
	assert.Nil(t, action.IsValid(), errorMessage)
	action = "READ"
	assert.Nil(t, action.IsValid(), errorMessage)
	action = "UPDATE"
	assert.Nil(t, action.IsValid(), errorMessage)
}
