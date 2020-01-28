package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryRead(t *testing.T) {
	client := NewClient(getDefaultConfig())

	name := "maven-central"
	repoType := "proxy"
	format := "maven2"

	repo, err := client.RepositoryRead(name, format, repoType)
	assert.Nil(t, err)
	assert.NotNil(t, repo)

	if repo != nil {
		assert.Equal(t, name, repo.Name)
	}
}
