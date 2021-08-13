package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDAPList(t *testing.T) {
	client := getTestClient()

	ldapServer, err := client.LDAPList()
	assert.Nil(t, err)
	assert.NotNil(t, ldapServer)
}

func TestLDAP(t *testing.T) {
	client := getTestClient()

	// https://hub.docker.com/r/mwaeckerlin/openldap/
	ldap := LDAP{
		AuthPassword:              "1234567890",
		AuthSchema:                "SIMPLE",
		AuthUserName:              "admin",
		ConnectionTimeoutSeconds:  uint(1),
		GroupType:                 "STATIC",
		Host:                      "127.0.0.1",
		Name:                      "ci-test",
		Port:                      389,
		Protocol:                  "LDAP",
		SearchBase:                "dc=example,dc=com",
		UserEmailAddressAttribute: "mail",
		UserIDAttribute:           "uid",
		UserObjectClass:           "inetOrgPerson",
		UserRealNameAttribute:     "cn",
	}
	err := client.LDAPCreate(ldap)
	assert.Nil(t, err)

	if err == nil {
		createdLDAP, err := client.LDAPRead(ldap.Name)
		assert.Nil(t, err)
		assert.NotNil(t, createdLDAP)
		assert.NotNil(t, createdLDAP.ID)
		assert.Equal(t, ldap.AuthSchema, createdLDAP.AuthSchema)
		assert.Equal(t, ldap.AuthUserName, createdLDAP.AuthUserName)
		assert.Equal(t, ldap.Host, createdLDAP.Host)
		assert.Equal(t, ldap.Port, createdLDAP.Port)
		assert.Equal(t, ldap.ConnectionTimeoutSeconds, createdLDAP.ConnectionTimeoutSeconds)
		// GroupType is not returned :-/
		// assert.Equal(t, ldap.GroupType, createdLDAP.GroupType)
		assert.Equal(t, ldap.Name, createdLDAP.Name)
		assert.Equal(t, ldap.Protocol, createdLDAP.Protocol)
		assert.Equal(t, ldap.SearchBase, createdLDAP.SearchBase)
		assert.Equal(t, ldap.UserEmailAddressAttribute, createdLDAP.UserEmailAddressAttribute)
		assert.Equal(t, ldap.UserIDAttribute, createdLDAP.UserIDAttribute)
		assert.Equal(t, ldap.UserObjectClass, createdLDAP.UserObjectClass)
		assert.Equal(t, ldap.UserRealNameAttribute, createdLDAP.UserRealNameAttribute)

		createdLDAP.Host = "127.0.0.2"
		createdLDAP.ID = ""
		// As GroupType is not returned while read, it needs to be set again
		createdLDAP.GroupType = "DYNAMIC"
		// As AuthPassword is not returned while read, it needs to be set again
		createdLDAP.AuthPassword = ldap.AuthPassword
		err = client.LDAPUpdate(ldap.Name, *createdLDAP)
		assert.Nil(t, err)

		updatedLDAP, err := client.LDAPRead(createdLDAP.Name)
		assert.Nil(t, err)
		assert.NotNil(t, updatedLDAP)
		assert.Equal(t, createdLDAP.Host, updatedLDAP.Host)

		err = client.LDAPDelete(ldap.Name)
		assert.Nil(t, err)

		deletedLDAP, err := client.LDAPRead(ldap.Name)
		// If the server was delete we get 404 and therefore expect an error
		assert.NotNil(t, err)
		assert.Nil(t, deletedLDAP)
	}
}
