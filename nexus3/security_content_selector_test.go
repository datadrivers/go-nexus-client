package nexus3

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSecurityContentSelector(t *testing.T) {
	client := getTestClient()
	name := "test-content-selector"
	createdContentSelector := testContentSelector(name)

	err := client.Security.ContentSelector.Create(*createdContentSelector)
	assert.Nil(t, err)

	listContentSelector, err := client.Security.ContentSelector.List()
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(listContentSelector))

	readContentSelector, err := client.Security.ContentSelector.Get(name)
	assert.Equal(t, createdContentSelector, readContentSelector)
	assert.Nil(t, err)

	updatedContentSelector := testContentSelector(name)

	err = client.Security.ContentSelector.Update(name, *updatedContentSelector)
	assert.Nil(t, err)

	readContentSelector, err = client.Security.ContentSelector.Get(name)
	assert.Equal(t, updatedContentSelector, readContentSelector)
	assert.Nil(t, err)

	err = client.Security.ContentSelector.Delete(name)
	assert.Nil(t, err)

	readContentSelector, err = client.Security.ContentSelector.Get(name)
	assert.Nil(t, readContentSelector)
	assert.Nil(t, err)
}

func testContentSelector(name string) *SecurityContentSelector {
	time := time.Now().Unix()
	return &SecurityContentSelector{
		Name:        name,
		Description: fmt.Sprintf("Go client content selector %d", time),
		Expression:  fmt.Sprintf("path == \"%d/\"", time),
	}
}
