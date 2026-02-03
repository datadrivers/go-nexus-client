package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/tools"
)

var (
	testClient *Client = nil
)

func getTestClient() *Client {
	if testClient != nil {
		return testClient
	}
	return NewClient(getDefaultConfig())
}

func getDefaultConfig() Config {
	return Config{
		Insecure: tools.GetEnv("NEXUS_INSECURE_SKIP_VERIFY", true).(bool),
		Password: tools.GetEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      tools.GetEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: tools.GetEnv("NEXUS_USRNAME", "admin").(string),
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(getDefaultConfig())

	assert.NotNil(t, c, "NewClient() must not return nil")
}

func TestContentType(t *testing.T) {
	c := getTestClient()

	c.ContentTypeJSON()
	assert.Equal(t, c.ContentType(), ContentTypeApplicationJSON)

	c.ContentTypeTextPlain()
	assert.Equal(t, c.ContentType(), ContentTypeTextPlain)
}
