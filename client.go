package client

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
)

// Client represents the Nexus API Client interface
type Client interface {
	CertificateList() (*[]Certificate, error)
	CertificateGet(*CertificateRequest) (*Certificate, error)
	CertificateCreate(*Certificate) error
	CertificateDelete(string) error
	ContentSelectorCreate(ContentSelector) error
	ContentSelectorRead(string) (*ContentSelector, error)
	ContentSelectorUpdate(string, ContentSelector) error
	ContentSelectorDelete(string) error
	ContentType() string
	ContentTypeTextPlain()
	ContentTypeJSON()
	BlobstoreCreate(Blobstore) error
	BlobstoreRead(string) (*Blobstore, error)
	BlobstoreUpdate(string, Blobstore) error
	BlobstoreDelete(string) error
	Privileges() ([]Privilege, error)
	PrivilegeCreate(Privilege) error
	PrivilegeRead(string) (*Privilege, error)
	PrivilegeUpdate(string, Privilege) error
	PrivilegeDelete(string) error
	RepositoryCreate(Repository) error
	RepositoryRead(string) (*Repository, error)
	RepositoryUpdate(string, Repository) error
	RepositoryDelete(string) error
	RoleCreate(Role) error
	RoleRead(string) (*Role, error)
	RoleUpdate(string, Role) error
	RoleDelete(string) error
	UserCreate(User) error
	UserRead(string) (*User, error)
	UserUpdate(string, User) error
	UserDelete(string) error
	UserChangePassword(string, string) error
	ScriptLists() ([]Script, error)
	ScriptRead(string) (*Script, error)
	ScriptCreate(*Script) error
	ScriptUpdate(*Script) error
	ScriptDelete(string) error
	ScriptRun(string) error
}

type client struct {
	config      Config
	contentType string
	httpClient  *http.Client
}

// NewClient returns an instance of client that implements the Client interface
func NewClient(config Config) Client {
	return &client{
		config:      config,
		contentType: ContentTypeApplicationJSON,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: config.Insecure,
				},
			},
		},
	}
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
