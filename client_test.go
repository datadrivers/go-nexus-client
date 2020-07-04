package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testClient Client = nil
)

func getTestClient() Client {
	if testClient != nil {
		return testClient
	}
	return NewClient(getDefaultConfig())
}

func getDefaultConfig() Config {
	return Config{
		Insecure: getEnv("NEXUS_INSECURE_SKIP_VERIFY", true).(bool),
		Password: getEnv("NEXUS_PASSWORD", "admin123").(string),
		URL:      getEnv("NEXUS_URL", "http://127.0.0.1:8081").(string),
		Username: getEnv("NEXUS_USRNAME", "admin").(string),
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

func TestPut(t *testing.T) {

}

func TestGet(t *testing.T) {

}

func TestPost(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
