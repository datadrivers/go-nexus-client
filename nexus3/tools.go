package nexus3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"

	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

func jsonMarshalInterfaceToIOReader(data interface{}) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("could not marshal data: %v", err)
	}

	return bytes.NewReader(b), nil
}

func getEnv(key string, fallback interface{}) interface{} {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
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
	request := &security.CertificateRequest{data.Hostname(), port}
	return request, nil
}

func getIntPointer(number int) *int {
	return &number
}

func getStringPointer(s string) *string {
	return &s
}

func getBoolPointer(b bool) *bool {
	return &b
}
