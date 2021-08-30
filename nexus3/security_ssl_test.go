package nexus3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecuritySSLCertificate(t *testing.T) {
	client := NewClient(getDefaultConfig())
	request, err := NewCertificateRequest("https://www.google.com")
	assert.Nil(t, err)

	certificate, err := client.Security.SSL.GetCertificate(request)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	err = client.Security.SSL.AddCertificate(certificate)
	assert.Nil(t, err)

	err = client.Security.SSL.RemoveCertificate(certificate.Id)
	assert.Nil(t, err)
}
