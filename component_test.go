package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComponentReadRead(t *testing.T) {
	client := getTestClient()

	id := ""

	component, err := client.ComponentRead(id)
	assert.Nil(t, err)
	assert.NotNil(t, component)

	if component != nil {
		assert.Equal(t, id, component.ID)
		assert.Equal(t, "maven-central", component.Repository)
	}
}
