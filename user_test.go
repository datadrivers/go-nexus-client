package client

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONUnmarshalUsers(t *testing.T) {
	testUsers := []User{*testUser("test-json-unmarshal-users")}

	testData, err := json.Marshal(testUsers)
	if err != nil {
		t.Fatalf("could not marshal testUsers: %v", err)
	}

	users, err := jsonUnmarshalUsers(testData)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func testUser(id string) *User {
	return &User{
		UserID:       id,
		FirstName:    "Test Firstname",
		LastName:     "Test Lastname",
		EmailAddress: "test-user@example.org",
		Password:     "abc123",
		Roles:        []string{"nx-admin"},
		Status:       "active",
	}
}

func TestUserCreateReadUpdateDelete(t *testing.T) {
	client := getTestClient()
	testUser := testUser("test-user-create-read-update-delete")

	err := client.UserCreate(*testUser)
	assert.Nil(t, err)

	user, err := client.UserRead(testUser.UserID)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, testUser.UserID, user.UserID)
	assert.Equal(t, testUser.FirstName, user.FirstName)
	assert.Equal(t, testUser.LastName, user.LastName)
	assert.Equal(t, testUser.EmailAddress, user.EmailAddress)
	assert.Equal(t, testUser.Status, user.Status)

	updatedUser := user
	updatedUser.FirstName = "changed"
	updatedUser.LastName = "changed"
	updatedUser.EmailAddress = "changed@example.com"

	err = client.UserUpdate(testUser.UserID, *updatedUser)
	assert.Nil(t, err)

	err = client.UserDelete(updatedUser.UserID)
	assert.Nil(t, err)
}

func TestUserCreate(t *testing.T) {
	client := getTestClient()
	testUser := testUser("test-user-create")

	err := client.UserCreate(*testUser)
	assert.Nil(t, err)

	err = client.UserDelete(testUser.UserID)
	assert.Nil(t, err)
}

func TestUserRead(t *testing.T) {
	client := getTestClient()

	user, err := client.UserRead("admin")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "admin", user.UserID)
	assert.Equal(t, "Administrator", user.FirstName)
	assert.Equal(t, "User", user.LastName)

}

func TestUserDeleteCurrentlySignedInUser(t *testing.T) {
	client := getTestClient()

	err := client.UserDelete(getDefaultConfig().Username)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Can not delete currently signed in user")
}
