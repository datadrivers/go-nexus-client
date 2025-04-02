package security

import (
	"encoding/json"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func TestJSONUnmarshalUsers(t *testing.T) {
	testUsers := []security.User{*testUser("test-json-unmarshal-users")}

	testData, err := json.Marshal(testUsers)
	if err != nil {
		t.Fatalf("could not marshal testUsers: %v", err)
	}

	users, err := jsonUnmarshalUsers(testData)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func testUser(id string) *security.User {
	return &security.User{
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
	service := getTestService()
	testUser := testUser("test-user-create-read-update-delete")

	err := service.User.Create(*testUser)
	assert.Nil(t, err)

	userSource := "default"

	user, err := service.User.Get(testUser.UserID, &userSource)
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

	err = service.User.Update(testUser.UserID, *updatedUser)
	assert.Nil(t, err)

	err = service.User.Delete(updatedUser.UserID)
	assert.Nil(t, err)
}

func TestUserCreate(t *testing.T) {
	service := getTestService()
	testUser := testUser("test-user-create")

	err := service.User.Create(*testUser)
	assert.Nil(t, err)

	err = service.User.Delete(testUser.UserID)
	assert.Nil(t, err)
}

func TestUserRead(t *testing.T) {
	service := getTestService()

	userSource := "default"
	user, err := service.User.Get("admin", &userSource)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "admin", user.UserID)
	assert.Equal(t, "Administrator", user.FirstName)
	assert.Equal(t, "User", user.LastName)

}

func TestUserDeleteCurrentlySignedInUser(t *testing.T) {
	service := getTestService()

	err := service.User.Delete(getDefaultConfig().Username)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Can not delete currently signed in user")
}

func TestListUsers(t *testing.T) {
	service := getTestService()

	users, err := service.User.List(nil)
	assert.Nil(t, err)
	// There are at least 2 users in a fresh Nexus installation
	assert.Equal(t, 2, len(users))

}
