package client

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONUnmarshalUsers(t *testing.T) {
	testUsers := []User{
		{
			UserID:       "user1",
			FirstName:    "Firstname",
			LastName:     "Lastname",
			EmailAddress: "user@example.org",
			Password:     "abs123",
			Status:       "active",
		},
		{},
	}

	testData, err := json.Marshal(testUsers)
	if err != nil {
		t.Fatalf("could not marshal testUsers: %v", err)
	}

	users, err := jsonUnmarshalUsers(testData)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func TestUserCreate(t *testing.T) {

}

func TestUserUpdate(t *testing.T) {

}

func TestUserChangePassword(t *testing.T) {

}

func TestUserRead(t *testing.T) {

}

func TestUserDelete(t *testing.T) {

}
