package nexus3

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// ContentTypeApplicationJSON ...
	ContentTypeApplicationJSON = "application/json"
	// ContentTypeTextPlain ...
	ContentTypeTextPlain = "text/plain"
	basePath             = "service/rest/"
)

type service struct {
	client *client
}

type client struct {
	config      Config
	contentType string
	httpClient  *http.Client

	// API Services
	BlobStore *BlobStoreService
	Security  *SecurityService
}

// NewClient returns an instance of client that implements the Client interface
func NewClient(config Config) *client {
	c := &client{
		config:      config,
		contentType: ContentTypeApplicationJSON,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: config.Insecure,
				},
			},
		},
	}
	c.BlobStore = NewBlobStoreService(c)
	c.Security = NewSecurityService(c)
	return c
}

func (c *client) setContentType(s string) {
	c.contentType = s
}

// ContentType returns the current configured HTTP content type
func (c *client) ContentType() string {
	return c.contentType
}

// ContentTypJSON configures the content type for future requests to be 'application/json'
func (c *client) ContentTypeJSON() {
	c.setContentType(ContentTypeApplicationJSON)
}

// ContentTypTestPlain configures the content typ for future requests to be 'test/plain'
func (c *client) ContentTypeTextPlain() {
	c.setContentType(ContentTypeTextPlain)
}

func (c *client) NewRequest(method string, endpoint string, body io.Reader) (req *http.Request, err error) {
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

func (c *client) execute(method string, endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
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

func (c *client) Get(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	return c.execute(http.MethodGet, endpoint, payload)
}

func (c *client) Post(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	return c.execute(http.MethodPost, endpoint, payload)
}

func (c *client) Put(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	return c.execute(http.MethodPut, endpoint, payload)
}

func (c *client) Delete(endpoint string) ([]byte, *http.Response, error) {
	return c.execute(http.MethodDelete, endpoint, nil)
}
