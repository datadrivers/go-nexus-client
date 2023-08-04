package nexus3

import (
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema"
	"github.com/stretchr/testify/assert"
)

func testMailConfig(enabled *bool, host string, port int, username *string, password *string, fromAddress string, subjectPrefix *string) *schema.MailConfig {
	b := new(bool)
	*b = false

	return &schema.MailConfig{
		Enabled:                       enabled,
		Host:                          host,
		Port:                          port,
		Username:                      username,
		Password:                      password,
		FromAddress:                   fromAddress,
		SubjectPrefix:                 subjectPrefix,
		StartTlsEnabled:               b,
		StartTlsRequired:              b,
		SslOnConnectEnabled:           b,
		SslServerIdentityCheckEnabled: b,
		NexusTrustStoreEnabled:        b,
	}
}

func TestMailConfigCreateReadUpdateDelete(t *testing.T) {
	client := getTestClient()

	enabled := true
	username := "uname"
	usernameUpdated := "username"
	password := "secret"
	subjectPrefix := "prefix"

	// Create
	err := client.MailConfig.Create(testMailConfig(&enabled, "example.org", 42, &username, &password, "sender@example.org", &subjectPrefix))
	assert.Nil(t, err)

	// Read
	readMailConfig, err := client.MailConfig.Get()
	assert.Nil(t, err)
	assert.Equal(t, readMailConfig, testMailConfig(&enabled, "example.org", 42, &username, nil, "sender@example.org", &subjectPrefix))

	// Update
	err = client.MailConfig.Update(testMailConfig(&enabled, "example.org", 42, &usernameUpdated, nil, "sender@example.org", &subjectPrefix))
	assert.Nil(t, err)

	// Check updated value
	readMailConfig, _ = client.MailConfig.Get()
	assert.Equal(t, readMailConfig, testMailConfig(&enabled, "example.org", 42, &usernameUpdated, nil, "sender@example.org", &subjectPrefix))

	// Delete
	err = client.MailConfig.Delete()
	assert.Nil(t, err)
}
