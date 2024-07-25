package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	// ContentTypeApplicationJSON ...
	ContentTypeApplicationJSON = "application/json"
	// ContentTypeTextPlain ...
	ContentTypeTextPlain = "text/plain"
	//
	BasePath = "service/rest/"
)

type Client struct {
	config      Config
	contentType string
	httpClient  *http.Client
}

// NewClient returns an instance of client that implements the Client interface
func NewClient(config Config) *Client {
	// Set default timeout value if not provided
	if config.Timeout == nil {
		defaultTimeout := 30
		config.Timeout = &defaultTimeout
	}

	var caCertPool *x509.CertPool
	if *config.RootCAPath != "" {
		caCert, err := os.ReadFile(*config.RootCAPath)
		// Backwards because we need to return a client and haven't got a logger
		if err == nil {
			caCertPool, err := x509.SystemCertPool()
			if err != nil {
				caCertPool = x509.NewCertPool()
			}
			caCertPool.AppendCertsFromPEM(caCert)
		}
	}

	var cert tls.Certificate
	if *config.ClientCertificatePath != "" && *config.ClientKeyPath != "" {
		// Load client PEM mTLS certificate
		cert, _ = tls.LoadX509KeyPair(*config.ClientCertificatePath, *config.ClientKeyPath)
	}

	return &Client{
		config:      config,
		contentType: ContentTypeApplicationJSON,
		httpClient: &http.Client{
			Timeout: time.Duration(*config.Timeout) * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: config.Insecure,
					RootCAs:            caCertPool,
					Certificates:       []tls.Certificate{cert},
				},
			},
		},
	}
}

func (c *Client) setContentType(s string) {
	c.contentType = s
}

// ContentType returns the current configured HTTP content type
func (c *Client) ContentType() string {
	return c.contentType
}

// ContentTypJSON configures the content type for future requests to be 'application/json'
func (c *Client) ContentTypeJSON() {
	c.setContentType(ContentTypeApplicationJSON)
}

// ContentTypTestPlain configures the content typ for future requests to be 'test/plain'
func (c *Client) ContentTypeTextPlain() {
	c.setContentType(ContentTypeTextPlain)
}

func (c *Client) NewRequest(method string, endpoint string, body io.Reader) (req *http.Request, err error) {
	url := fmt.Sprintf("%s/%s", c.config.URL, endpoint)
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}

	req.SetBasicAuth(c.config.Username, c.config.Password)
	req.Header.Set("Content-Type", c.contentType)
	req.Header.Set("Accept", ContentTypeApplicationJSON)

	return req, nil
}

func (c *Client) execute(method string, endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	req, err := c.NewRequest(method, endpoint, payload)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, resp, err
}

func (c *Client) Get(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	return c.execute(http.MethodGet, endpoint, payload)
}

func (c *Client) Post(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	return c.execute(http.MethodPost, endpoint, payload)
}

func (c *Client) Put(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	return c.execute(http.MethodPut, endpoint, payload)
}

func (c *Client) Delete(endpoint string) ([]byte, *http.Response, error) {
	return c.execute(http.MethodDelete, endpoint, nil)
}
