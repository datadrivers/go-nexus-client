package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCertificate(t *testing.T) {
	client := NewClient(getDefaultConfig())
	request, err := NewCertificateRequest("https://www.google.com")
	assert.Nil(t, err)

	certificate, err := client.CertificateGet(request)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)

	err = client.CertificateCreate(certificate)
	assert.Nil(t, err)

	err = client.CertificateDelete(certificate.Id)
	assert.Nil(t, err)
}
