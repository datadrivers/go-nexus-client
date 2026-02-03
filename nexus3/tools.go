package nexus3

import (
	"net/url"
	"strconv"

	"github.com/williamt1997/go-nexus-client/nexus3/schema/security"
)

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
	request := &security.CertificateRequest{data.Hostname(), port}
	return request, nil
}
