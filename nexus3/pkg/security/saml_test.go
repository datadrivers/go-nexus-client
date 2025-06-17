package security

import (
	"io/ioutil"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
	"github.com/stretchr/testify/assert"
)

func TestSAML(t *testing.T) {
	if tools.GetEnv("SKIP_PRO_TESTS", "false") == "true" {
		t.Skip("Skipping Nexus Pro tests")
	}
	service := getTestService()

	// https://samltest.id/saml/idp
	// configuration is not fully validated until an authentication via SSO is attempted
	// we will use invalid values apart from the IdP metadata to verify the configuration is applied

	dat, err := ioutil.ReadFile("./testfiles/saml-testconfig.xml")
	assert.Nil(t, err)

	validateAssertionSignature := false
	validateResponseSignature := false
	firstNameAttribute := "firstName2"
	lastNameAttribute := "lastName2"
	emailAttribute := "email2"
	groupsAttribute := "groups2"

	saml := security.SAML{
		IdpMetadata:                string(dat),
		EntityId:                   "http://localhost:8081/service/rest/v1/security/saml/metadata",
		ValidateAssertionSignature: &validateAssertionSignature,
		ValidateResponseSignature:  &validateResponseSignature,
		UsernameAttribute:          "username2",
		FirstNameAttribute:         &firstNameAttribute,
		LastNameAttribute:          &lastNameAttribute,
		EmailAttribute:             &emailAttribute,
		GroupsAttribute:            &groupsAttribute,
	}
	err = service.SAML.Apply(saml)
	assert.Nil(t, err)

	createdSAML, err := service.SAML.Read()
	assert.Nil(t, err)
	assert.NotNil(t, createdSAML)
	assert.NotNil(t, createdSAML.IdpMetadata)
	assert.NotNil(t, createdSAML.UsernameAttribute)
	assert.Equal(t, saml.IdpMetadata, createdSAML.IdpMetadata)
	assert.Equal(t, saml.EntityId, createdSAML.EntityId)
	assert.Equal(t, saml.ValidateAssertionSignature, createdSAML.ValidateAssertionSignature)
	assert.Equal(t, saml.ValidateResponseSignature, createdSAML.ValidateResponseSignature)
	assert.Equal(t, saml.UsernameAttribute, createdSAML.UsernameAttribute)
	assert.Equal(t, saml.FirstNameAttribute, createdSAML.FirstNameAttribute)
	assert.Equal(t, saml.LastNameAttribute, createdSAML.LastNameAttribute)
	assert.Equal(t, saml.EmailAttribute, createdSAML.EmailAttribute)
	assert.Equal(t, saml.GroupsAttribute, createdSAML.GroupsAttribute)

	createdSAML.UsernameAttribute = "username"
	err = service.SAML.Apply(*createdSAML)
	assert.Nil(t, err)

	updatedSAML, err := service.SAML.Read()
	assert.Nil(t, err)
	assert.NotNil(t, updatedSAML)
	assert.Equal(t, createdSAML.UsernameAttribute, updatedSAML.UsernameAttribute)

	err = service.SAML.Delete()
	assert.Nil(t, err)

	deletedSAML, err := service.SAML.Read()
	// If the configuration was delete we get 404 and therefore expect an error
	assert.NotNil(t, err)
	assert.Nil(t, deletedSAML)
}
