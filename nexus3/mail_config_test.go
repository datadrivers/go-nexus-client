package nexus3

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema"
	"github.com/stretchr/testify/assert"
)

// When reading the mail config Nexus does not return the password but this value
var emptyPasswordValue = ""

func testMailConfig(enabled bool, host string, port int, username string, password string, fromAddress string, subjectPrefix string) *schema.MailConfig {
	return &schema.MailConfig{
		Enabled:                       enabled,
		Host:                          host,
		Port:                          port,
		Username:                      username,
		Password:                      password,
		FromAddress:                   fromAddress,
		SubjectPrefix:                 subjectPrefix,
		StartTlsEnabled:               false,
		StartTlsRequired:              false,
		SslOnConnectEnabled:           false,
		SslServerIdentityCheckEnabled: false,
		NexusTrustStoreEnabled:        false,
	}
}

func TestMailConfigCreateReadUpdateDelete(t *testing.T) {
	client := getTestClient()

	// Create
	err := client.MailConfig.Create(testMailConfig(true, "example.org", 42, "uname", "secret", "sender@example.org", "SubjectPrefix"))
	assert.Nil(t, err)

	// Read
	readMailConfig, err := client.MailConfig.Get()
	assert.Nil(t, err)
	assert.Equal(t, readMailConfig, testMailConfig(true, "example.org", 42, "uname", emptyPasswordValue, "sender@example.org", "SubjectPrefix"))

	// Update
	err = client.MailConfig.Update(testMailConfig(true, "example.org", 42, "username", emptyPasswordValue, "sender@example.org", "SubjectPrefix"))
	assert.Nil(t, err)

	// Check updated value
	readMailConfig, _ = client.MailConfig.Get()
	assert.Equal(t, readMailConfig, testMailConfig(true, "example.org", 42, "username", emptyPasswordValue, "sender@example.org", "SubjectPrefix"))

	// Delete
	err = client.MailConfig.Delete()
	assert.Nil(t, err)
}
