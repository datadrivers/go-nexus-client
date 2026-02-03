package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/security"
)

func TestSecurityLDAPList(t *testing.T) {
	service := getTestService()

	ldapServer, err := service.LDAP.List()
	assert.Nil(t, err)
	assert.NotNil(t, ldapServer)
}

func TestSecurityLDAP(t *testing.T) {
	service := getTestService()

	// https://hub.docker.com/r/mwaeckerlin/openldap/
	ldap := security.LDAP{
		AuthPassword:              "1234567890",
		AuthSchema:                "SIMPLE",
		AuthUserName:              "admin",
		ConnectionTimeoutSeconds:  int32(1),
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
	err := service.LDAP.Create(ldap)
	assert.Nil(t, err)

	createdLDAP, err := service.LDAP.Get(ldap.Name)
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
	err = service.LDAP.Update(ldap.Name, *createdLDAP)
	assert.Nil(t, err)

	updatedLDAP, err := service.LDAP.Get(createdLDAP.Name)
	assert.Nil(t, err)
	assert.NotNil(t, updatedLDAP)
	assert.Equal(t, createdLDAP.Host, updatedLDAP.Host)

	err = service.LDAP.Delete(ldap.Name)
	assert.Nil(t, err)

	deletedLDAP, err := service.LDAP.Get(ldap.Name)
	// If the server was delete we get 404 and therefore expect an error
	assert.NotNil(t, err)
	assert.Nil(t, deletedLDAP)
}
