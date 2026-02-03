package security

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/security"
)

const (
	securitySslAPIEndpoint = securityAPIEndpoint + "/ssl"
)

type SecuritySSLService client.Service

func NewSecuritySSLService(c *client.Client) *SecuritySSLService {

	s := &SecuritySSLService{
		Client: c,
	}
	return s
}

func jsonUnmarshalCertificate(data []byte) (*security.SSLCertificate, error) {
	var certificate = security.SSLCertificate{}
	if err := json.Unmarshal(data, &certificate); err != nil {
		return nil, fmt.Errorf("could not unmarshal certificate: %v", err)
	}
	return &certificate, nil
}

func jsonUnmarshalCertificateList(data []byte) (*[]security.SSLCertificate, error) {
	var certificates []security.SSLCertificate
	if err := json.Unmarshal(data, &certificates); err != nil {
		return nil, fmt.Errorf("could not unmarshal certificates: %v", err)
	}
	return &certificates, nil
}

// Add a certificate to the trust store
func (s *SecuritySSLService) AddCertificate(certificate *security.SSLCertificate) error {
	data := strings.NewReader(certificate.Pem)

	body, resp, err := s.Client.Post(fmt.Sprintf("%s/truststore", securitySslAPIEndpoint), data)
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
	body, resp, err := s.Client.Delete(fmt.Sprintf("%s/truststore/%s", securitySslAPIEndpoint, url.QueryEscape(id)))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete certificate '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

// Retrieve a list of certificates added to the trust store
func (s *SecuritySSLService) ListCertificates() (*[]security.SSLCertificate, error) {
	body, resp, err := s.Client.Get(fmt.Sprintf("%s/truststore", securitySslAPIEndpoint), nil)
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
func (s *SecuritySSLService) GetCertificate(params *security.CertificateRequest) (*security.SSLCertificate, error) {
	values, _ := query.Values(&params)

	body, resp, err := s.Client.Get(fmt.Sprintf("%s?%s", securitySslAPIEndpoint, values.Encode()), nil)

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

func NewCertificateRequest(proxyUrl string) (*security.CertificateRequest, error) {
	data, err := url.Parse(proxyUrl)
	if err != nil {
		return nil, err
	}
	port := 443
	if data.Port() != "" {
		port, err = strconv.Atoi(data.Port())
		if err != nil {
			return nil, err
		}
	}
	request := &security.CertificateRequest{Host: data.Hostname(), Port: port}
	return request, nil
}
