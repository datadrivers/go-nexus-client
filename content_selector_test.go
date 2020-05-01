package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestContentSelectorCreateReadUpdateDelete(t *testing.T) {
	client := NewClient(getDefaultConfig())
	name := "test-content-selector"
	createdContentSelector := testContentSelector(name)

	err := client.ContentSelectorCreate(*createdContentSelector)
	assert.Nil(t, err)

	readContentSelector, err := client.ContentSelectorRead(name)
	assert.Equal(t, createdContentSelector, readContentSelector)
	assert.Nil(t, err)

	updatedContentSelector := testContentSelector(name)

	err = client.ContentSelectorUpdate(name, *updatedContentSelector)
	assert.Nil(t, err)

	readContentSelector, err = client.ContentSelectorRead(name)
	assert.Equal(t, updatedContentSelector, readContentSelector)
	assert.Nil(t, err)

	err = client.ContentSelectorDelete(name)
	assert.Nil(t, err)

	readContentSelector, err = client.ContentSelectorRead(name)
	assert.Nil(t, readContentSelector)
	assert.Nil(t, err)
}

func testContentSelector(name string) *ContentSelector {
	time := time.Now().Unix()
	return &ContentSelector{
		Name:        name,
		Description: fmt.Sprintf("Go client content selector %d", time),
		Expression:  fmt.Sprintf("path == \"%d/\"", time),
	}
}
