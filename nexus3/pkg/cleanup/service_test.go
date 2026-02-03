package cleanup_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/cleanup"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/tools"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/cleanuppolicies"
)

// https://help.sonatype.com/en/cleanup-policies-api.html

var (
	testClient *client.Client = nil
)

func getTestClient() *client.Client {
	if testClient != nil {
		return testClient
	}
	return client.NewClient(getDefaultConfig())
}

func getTestService() *cleanup.CleanupPolicyService {
	return cleanup.NewCleanupPolicyService(getTestClient())
}

func getDefaultConfig() client.Config {
	return client.Config{
		Insecure: tools.GetEnv("NEXUS_INSECURE_SKIP_VERIFY", true).(bool),
		Password: tools.GetEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      tools.GetEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: tools.GetEnv("NEXUS_USRNAME", "admin").(string),
	}
}

func TestNewCleanupService(t *testing.T) {
	s := getTestService()

	assert.NotNil(t, s, "NewCleanupService() must not return nil")
}

func TestCreateCleanupPolicy(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}

	s := getTestService()

	policy := &cleanuppolicies.CleanupPolicy{
		Notes:              tools.GetStringPointer("Test"),
		CriteriaAssetRegex: tools.GetStringPointer("/*"),
		Name:               "Test",
		Format:             cleanuppolicies.RepositoryFormatGo,
	}

	policy2 := &cleanuppolicies.CleanupPolicy{
		Notes:                  tools.GetStringPointer("Test2"),
		CriteriaLastDownloaded: tools.GetIntPointer(1),
		Name:                   "Test2",
		Format:                 cleanuppolicies.RepositoryFormatGitLfs,
	}

	err := s.Create(policy)
	assert.Nil(t, err, "Create() must not return an error")

	err = s.Create(policy2)
	assert.Nil(t, err, "Second Create() must not return an error")

	policy, err = s.Get(policy.Name)
	assert.Nil(t, err, "Get() must not return an error")
	assert.Equal(t, policy.Name, "Test")

	policies, err := s.List()
	assert.Nil(t, err, "List() must not return an error")
	assert.Equal(t, len(policies), 2)

	policy.CriteriaLastBlobUpdated = tools.GetIntPointer(1)
	err = s.Update(policy)
	assert.Nil(t, err, "Update() must not return an error")
	assert.Equal(t, policy.CriteriaLastBlobUpdated, tools.GetIntPointer(1))

	err = s.Delete(policy.Name)
	assert.Nil(t, err, "Delete() must not return an error")

	err = s.Delete(policy2.Name)
	assert.Nil(t, err, "Second Delete() must not return an error")
}
