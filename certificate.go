package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	certificateAPIEndpoint = basePath + "v1/security/ssl"
)

type Certificate struct {
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

type CertificateRequest struct {
	Host string `url:"host"`
	Port int    `url:"port"`
}

func NewCertificateRequest(proxyUrl string) (*CertificateRequest, error) {
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
	request := &CertificateRequest{data.Hostname(), port}
	return request, nil
}

func jsonUnmarshalCertificate(data []byte) (*Certificate, error) {
	var certificate = Certificate{}
	if err := json.Unmarshal(data, &certificate); err != nil {
		return nil, fmt.Errorf("could not unmarshal certificate: %v", err)
	}
	return &certificate, nil
}

func jsonUnmarshalCertificateList(data []byte) (*[]Certificate, error) {
	var certificates []Certificate
	if err := json.Unmarshal(data, &certificates); err != nil {
		return nil, fmt.Errorf("could not unmarshal certificates: %v", err)
	}
	return &certificates, nil
}

func (c client) CertificateCreate(certificate *Certificate) error {
	data := strings.NewReader(certificate.Pem)

	body, resp, err := c.Post(fmt.Sprintf("%s/truststore", certificateAPIEndpoint), data)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create certificate '%s': HTTP: %d, %s", certificate.Id, resp.StatusCode, string(body))
	}
	return nil
}

func (c client) CertificateDelete(id string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/truststore/%s", certificateAPIEndpoint, url.QueryEscape(id)))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete certificate '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (c client) CertificateList() (*[]Certificate, error) {
	body, resp, err := c.Get(fmt.Sprintf("%s/truststore", certificateAPIEndpoint), nil)
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

func (c client) CertificateGet(params *CertificateRequest) (*Certificate, error) {
	values, _ := query.Values(&params)

	body, resp, err := c.Get(fmt.Sprintf("%s?%s", certificateAPIEndpoint, values.Encode()), nil)

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
