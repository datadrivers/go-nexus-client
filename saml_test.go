package client

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSAML(t *testing.T) {
	if getEnv("SKIP_PRO_TESTS", false).(bool) {
		t.Skip("Skipping Nexus Pro tests")
	}
	client := getTestClient()

	// https://samltest.id/saml/idp
	// configuration is not fully validated until an authentication via SSO is attempted
	// we will use invalid values apart from the IdP metadata to verify the configuration is applied

	dat, err := ioutil.ReadFile("./saml-testconfig.xml")
	assert.Nil(t, err)

	saml := SAML{
		IdpMetadata:                string(dat),
		EntityId:                   "http://localhost:8081/service/rest/v1/security/saml/metadata",
		ValidateAssertionSignature: false,
		ValidateResponseSignature:  false,
		UsernameAttribute:          "username2",
		FirstNameAttribute:         "firstName2",
		LastNameAttribute:          "lastName2",
		EmailAttribute:             "email2",
		GroupsAttribute:            "groups2",
	}
	err = client.SAMLApply(saml)
	assert.Nil(t, err)

	if err == nil {
		createdSAML, err := client.SAMLRead()
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
		err = client.SAMLApply(*createdSAML)
		assert.Nil(t, err)

		updatedSAML, err := client.SAMLRead()
		assert.Nil(t, err)
		assert.NotNil(t, updatedSAML)
		assert.Equal(t, createdSAML.UsernameAttribute, updatedSAML.UsernameAttribute)

		err = client.SAMLDelete()
		assert.Nil(t, err)

		deletedSAML, err := client.SAMLRead()
		// If the configuration was delete we get 404 and therefore expect an error
		assert.NotNil(t, err)
		assert.Nil(t, deletedSAML)
	}
}
