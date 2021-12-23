package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecuritySSLCertificate(t *testing.T) {
	service := getTestService()
	request, err := NewCertificateRequest("https://www.google.com")
	assert.Nil(t, err)

	certificate, err := service.SSL.GetCertificate(request)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	err = service.SSL.AddCertificate(certificate)
	assert.Nil(t, err)

	err = service.SSL.RemoveCertificate(certificate.Id)
	assert.Nil(t, err)
}
