package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrivilegeCreateReadUpdateDelete(t *testing.T) {
	client := NewClient(getDefaultConfig())
	name := "test-privilege"
	createdPrivilege := testPrivilege(name)

	err := client.PrivilegeCreate(*createdPrivilege)
	assert.Nil(t, err)

	readPrivilege, err := client.PrivilegeRead(name)
	assert.Equal(t, createdPrivilege, readPrivilege)
	assert.Nil(t, err)

	updatedPrivilege := testPrivilege(name)

	err = client.PrivilegeUpdate(name, *updatedPrivilege)
	assert.Nil(t, err)

	readPrivilege, err = client.PrivilegeRead(name)
	assert.Equal(t, updatedPrivilege, readPrivilege)
	assert.Nil(t, err)

	err = client.PrivilegeDelete(name)
	assert.Nil(t, err)

	readPrivilege, err = client.PrivilegeRead(name)
	assert.Nil(t, readPrivilege)
	assert.Nil(t, err)
}

func testPrivilege(name string) *Privilege {
	time := time.Now().Unix()
	return &Privilege{
		Actions:     []string{"READ"},
		Description: fmt.Sprintf("Go client privilege %d", time),
		Domain:      "*",
		Name:        name,
		Type:        "application",
	}
}
