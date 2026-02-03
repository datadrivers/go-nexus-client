package security

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/security"
)

func TestSecurityContentSelector(t *testing.T) {
	service := getTestService()
	name := "test-content-selector"
	createdContentSelector := testContentSelector(name)

	err := service.ContentSelector.Create(*createdContentSelector)
	assert.Nil(t, err)

	listContentSelector, err := service.ContentSelector.List()
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(listContentSelector))

	readContentSelector, err := service.ContentSelector.Get(name)
	assert.Equal(t, createdContentSelector, readContentSelector)
	assert.Nil(t, err)

	updatedContentSelector := testContentSelector(name)

	err = service.ContentSelector.Update(name, *updatedContentSelector)
	assert.Nil(t, err)

	readContentSelector, err = service.ContentSelector.Get(name)
	assert.Equal(t, updatedContentSelector, readContentSelector)
	assert.Nil(t, err)

	err = service.ContentSelector.Delete(name)
	assert.Nil(t, err)

	readContentSelector, err = service.ContentSelector.Get(name)
	assert.Nil(t, readContentSelector)
	assert.Nil(t, err)
}

func testContentSelector(name string) *security.ContentSelector {
	time := time.Now().Unix()
	return &security.ContentSelector{
		Name:        name,
		Description: fmt.Sprintf("Go client content selector %d", time),
		Expression:  fmt.Sprintf("path == \"%d/\"", time),
	}
}
