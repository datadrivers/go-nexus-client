package client

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getEnv(key, fallback string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		return fallback
	}
	return value
}

func getDefaultConfig() Config {
	return Config{
		URL:      getEnv("NEXUS_URL", "http://127.0.0.1:8080"),
		Username: getEnv("NEXUS_USRNAME", "admin"),
		Password: getEnv("NEXUS_PASSWORD", "admin123"),
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(getDefaultConfig())

	assert.NotNil(t, c, "NewClient() must not return nil")
}

func TestContentType(t *testing.T) {
	c := NewClient(getDefaultConfig())

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
