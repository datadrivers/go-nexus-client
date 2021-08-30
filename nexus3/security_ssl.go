package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	securitySslAPIEndpoint = securityAPIEndpoint + "/ssl"
)

type SecuritySSLService service

func NewSecuritySSLService(c *client) *SecuritySSLService {

	s := &SecuritySSLService{
		client: c,
	}
	return s
}

type SecuritySSLCertificate struct {
	Id                      string `json:"id"`
	Fingerprint             string `json:"fingerprint"`
	SerialNumber            string `json:"serialNumber"`
	IssuerCommonName        string `json:"issuerCommonName"`
	IssuerOrganization      string `json:"issuerOrganization"`
	IssuerOrganizationUnit  string `json:"issuerOrganizationalUnit"`
	SubjectCommonName       string `json:"subjectCommonName"`
	SubjectOrganization     string `json:"subjectOrganization"`
	SubjectOrganizationUnit string `json:"subjectOrganizationalUnit"`
	Pem                     string `json:"pem"`
	IssuedOn                int64  `json:"issuedOn"`
	ExpiresOn               int64  `json:"expiresOn"`
}

func jsonUnmarshalCertificate(data []byte) (*SecuritySSLCertificate, error) {
	var certificate = SecuritySSLCertificate{}
	if err := json.Unmarshal(data, &certificate); err != nil {
		return nil, fmt.Errorf("could not unmarshal certificate: %v", err)
	}
	return &certificate, nil
}

func jsonUnmarshalCertificateList(data []byte) (*[]SecuritySSLCertificate, error) {
	var certificates []SecuritySSLCertificate
	if err := json.Unmarshal(data, &certificates); err != nil {
		return nil, fmt.Errorf("could not unmarshal certificates: %v", err)
	}
	return &certificates, nil
}

// Add a certificate to the trust store
func (s *SecuritySSLService) AddCertificate(certificate *SecuritySSLCertificate) error {
	data := strings.NewReader(certificate.Pem)

	body, resp, err := s.client.Post(fmt.Sprintf("%s/truststore", securitySslAPIEndpoint), data)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create certificate '%s': HTTP: %d, %s", certificate.Id, resp.StatusCode, string(body))
	}
	return nil
}

// Remove a certificate in the trust store
func (s *SecuritySSLService) RemoveCertificate(id string) error {
	body, resp, err := s.client.Delete(fmt.Sprintf("%s/truststore/%s", securitySslAPIEndpoint, url.QueryEscape(id)))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete certificate '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

// Retrieve a list of certificates added to the trust store
func (s *SecuritySSLService) ListCertificates() (*[]SecuritySSLCertificate, error) {
	body, resp, err := s.client.Get(fmt.Sprintf("%s/truststore", securitySslAPIEndpoint), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get certificates: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	certificates, err := jsonUnmarshalCertificateList(body)

	if err != nil {
		return nil, err
	}

	return certificates, nil
}

// Get a certificate in the trust store
func (s *SecuritySSLService) GetCertificate(params *CertificateRequest) (*SecuritySSLCertificate, error) {
	values, _ := query.Values(&params)

	body, resp, err := s.client.Get(fmt.Sprintf("%s?%s", securitySslAPIEndpoint, values.Encode()), nil)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get certificate from '%s:%d': HTTP: %d, %s", params.Host, params.Port, resp.StatusCode, string(body))
	}

	certificate, err := jsonUnmarshalCertificate(body)
	if err != nil {
		return nil, err
	}

	return certificate, nil
}
