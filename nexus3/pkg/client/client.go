package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	// ContentTypeApplicationJSON ...
	ContentTypeApplicationJSON = "application/json"
	// ContentTypeTextPlain ...
	ContentTypeTextPlain = "text/plain"
	// BasePath ...
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
	if config.RootCAPath != nil && *config.RootCAPath != "" {
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
	if config.ClientCertificatePath != nil && *config.ClientCertificatePath != "" &&
		config.ClientKeyPath != nil && *config.ClientKeyPath != "" {
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

// ContentTypeJSON configures the content type for future requests to be 'application/json'
func (c *Client) ContentTypeJSON() {
	c.setContentType(ContentTypeApplicationJSON)
}

// ContentTypeTextPlain configures the content type for future requests to be 'test/plain'
func (c *Client) ContentTypeTextPlain() {
	c.setContentType(ContentTypeTextPlain)
}

func (c *Client) NewRequest(method string, endpoint string, body io.Reader) (req *http.Request, err error) {
	return c.NewRequestContext(context.Background(), method, endpoint, body)
}

func (c *Client) NewRequestContext(ctx context.Context, method string, endpoint string, body io.Reader) (req *http.Request, err error) {
	url := fmt.Sprintf("%s/%s", c.config.URL, endpoint)
	req, err = http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return req, err
	}

	req.SetBasicAuth(c.config.Username, c.config.Password)
	req.Header.Set("Content-Type", c.contentType)
	req.Header.Set("Accept", ContentTypeApplicationJSON)

	return req, nil
}

func (c *Client) execute(req *http.Request) ([]byte, *http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body, resp, err
}

func (c *Client) Get(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	if req, err := c.NewRequest(http.MethodGet, endpoint, payload); err != nil {
		return nil, nil, err
	} else {
		return c.execute(req)
	}
}

func (c *Client) Post(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	if req, err := c.NewRequest(http.MethodPost, endpoint, payload); err != nil {
		return nil, nil, err
	} else {
		return c.execute(req)
	}
}

func (c *Client) Put(endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	if req, err := c.NewRequest(http.MethodPut, endpoint, payload); err != nil {
		return nil, nil, err
	} else {
		return c.execute(req)
	}
}

func (c *Client) Delete(endpoint string) ([]byte, *http.Response, error) {
	if req, err := c.NewRequest(http.MethodDelete, endpoint, nil); err != nil {
		return nil, nil, err
	} else {
		return c.execute(req)
	}
}

func (c *Client) GetContext(ctx context.Context, endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	if req, err := c.NewRequestContext(ctx, http.MethodGet, endpoint, payload); err != nil {
		return nil, nil, err
	} else {
		return c.execute(req)
	}
}

func (c *Client) PostContext(ctx context.Context, endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	if req, err := c.NewRequestContext(ctx, http.MethodPost, endpoint, payload); err != nil {
		return nil, nil, err
	} else {
		return c.execute(req)
	}
}

func (c *Client) PutContext(ctx context.Context, endpoint string, payload io.Reader) ([]byte, *http.Response, error) {
	if req, err := c.NewRequestContext(ctx, http.MethodPut, endpoint, payload); err != nil {
		return nil, nil, err
	} else {
		return c.execute(req)
	}
}

func (c *Client) DeleteContext(ctx context.Context, endpoint string) ([]byte, *http.Response, error) {
	if req, err := c.NewRequestContext(ctx, http.MethodDelete, endpoint, nil); err != nil {
		return nil, nil, err
	} else {
		return c.execute(req)
	}
}
