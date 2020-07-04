package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCertificate(t *testing.T) {
	client := NewClient(getDefaultConfig())
	request, err := NewCertificateRequest("https://www.google.com")
	assert.Nil(t, err)
	fmt.Printf("%#v\n", request)

	certificate, err := client.CertificateGet(request)
	assert.Nil(t, err)
	assert.NotNil(t, certificate)
	fmt.Printf("%#v\n", certificate)

	err = client.CertificateCreate(certificate)
	assert.Nil(t, err)

	err = client.CertificateDelete(certificate.Id)
	assert.Nil(t, err)
}
