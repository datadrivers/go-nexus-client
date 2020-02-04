package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	repositoryAPIEndpoint = "service/rest/beta/repositories"

	RepositoryFormatApt    = "apt"
	RepositoryFormatBower  = "bower"
	RepositoryFormatDocker = "docker"
	RepositoryFormatMaven2 = "maven2"

	RepositoryTypeHosted = "hosted"
	RepositoryTypeGroup  = "group"
	RepositoryTypeProxy  = "proxy"
)

// Repository ...
type Repository struct {
	Format          string  `json:"format"`
	Name            string  `json:"name"`
	Online          bool    `json:"online"`
	RoutingRuleName *string `json:"routingRuleName,omitempty"`
	Type            string  `json:"type"`

	// Apt Repository data
	*RepositoryApt        `json:"apt,omitempty"`
	*RepositoryAptSigning `json:"aptSigning,omitempty"`

	// RepositoryCleanup data
	*RepositoryCleanup `json:"cleanup,omitempty"`

	// RepositoryBower data
	*RepositoryBower `json:"bower,omitempty"`

	// Docker Repository data
	*RepositoryDocker      `json:"docker,omitempty"`
	*RepositoryDockerProxy `json:"dockerProxy,omitempty"`

	// Group data
	*RepositoryGroup `json:"group,omitempty"`

	// HTTPClient
	*RepositoryHTTPClient `json:"httpClient,omitempty"`

	// Maven Reporitoy data
	*RepositoryMaven `json:"maven,omitempty"`

	// Cache data for proxy Repository
	*RepositoryNegativeCache `json:"negativeCache,omitempty"`

	// Proxy Repository data
	*RepositoryProxy `json:"proxy,omitempty"`

	// Repository storage data
	*RepositoryStorage `json:"storage"`
}

// RepositoryApt contains the data of an Apt Repository
type RepositoryApt struct {
	Distribution string `json:"distribution"`
}

// RepositoryAptSigning contains values for Apt signing
type RepositoryAptSigning struct {
	Keypair    string `json:"keypair"`
	Passphrase string `json:"passphrase"`
}

// RepositoryBower contains data of bower repositories
type RepositoryBower struct {
	RewritePackageUrls bool `json:"rewritePackageUrls"`
}

// RepositoryCleanup ...
type RepositoryCleanup struct {
	PolicyNames []string `json:"policyNames"`
}

// RepositoryDocker contains data of a Docker Repositoriy
type RepositoryDocker struct {
	ForceBasicAuth bool `json:"forceBasicAuth"`
	HTTPPort       *int `json:"httpPort"`
	HTTPSPort      *int `json:"httpsPort"`
	V1Enabled      bool `json:"v1Enabled"`
}

// RepositoryDockerProxy contains data of a Docker Proxy Repository
type RepositoryDockerProxy struct {
	IndexType string  `json:"indexType"`
	IndexURL  *string `json:"indexUrl,omitempty"`
}

// RepositoryGroup contains repository group configuration data
type RepositoryGroup struct {
	MemberNames []string `json:"memberNames,omitempty"`
}

// RepositoryHTTPClient ...
type RepositoryHTTPClient struct {
	Authentication RepositoryHTTPClientAuthentication `json:"authentication"`
	AutoBlock      bool                               `json:"autoBlock"`
	Blocked        bool                               `json:"blocked"`
	Connection     RepositoryHTTPClientConnection     `json:"connection"`
}

// RepositoryHTTPClientConnection ...
type RepositoryHTTPClientConnection struct {
	EnableCircularRedirects bool    `json:"enableCircularRedirects"`
	EnableCookies           bool    `json:"enableCookies"`
	Retries                 *int    `json:"retries"`
	Timeout                 *int    `json:"timeout"`
	UserAgentSuffic         *string `json:"userAgentSuffix"`
}

// RepositoryHTTPClientAuthentication ...
type RepositoryHTTPClientAuthentication struct {
	NTLMDomain string `json:"ntlmDomain"`
	NTLMHost   string `json:"ntlmHost"`
	Type       string `json:"type"`
	Username   string `json:"username"`
}

// RepositoryMaven contains additional data of maven repository
type RepositoryMaven struct {
	VersionPolicy string `json:"versionPolicy"`
	LayoutPolicy  string `json:"layoutPolicy"`
}

// RepositoryNegativeCache ...
type RepositoryNegativeCache struct {
	Enabled bool `json:"enabled"`
	TTL     int  `json:"timeToLive"`
}

// RepositoryProxy contains Proxy Repository data
type RepositoryProxy struct {
	ContentMaxAge  int    `json:"contentMaxAge"`
	MetadataMaxAge int    `json:"metadataMaxAge"`
	RemoteURL      string `json:"remoteUrl"`
}

// RepositoryStorage contains repository storage
type RepositoryStorage struct {
	BlobStoreName               string  `json:"blobStoreName"`
	StrictContentTypeValidation bool    `json:"strictContentTypeValidation"`
	WritePolicy                 *string `json:"writePolicy,omitempty"`
}

func jsonUnmarshalRepositories(data []byte) ([]Repository, error) {
	var repositories []Repository
	if err := json.Unmarshal(data, &repositories); err != nil {
		return nil, fmt.Errorf("could not unmarshal repositories: %v", err)
	}
	return repositories, nil
}

func (c client) RepositoryCreate(repo Repository) error {
	data, err := jsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}

	body, resp, err := c.Post(fmt.Sprintf("%s/%s/%s", repositoryAPIEndpoint, repo.Format, repo.Type), data)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (c client) RepositoryRead(id string) (*Repository, error) {
	body, resp, err := c.Get(fmt.Sprintf("%s", repositoryAPIEndpoint), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}

	repositories, err := jsonUnmarshalRepositories(body)
	if err != nil {
		return nil, err
	}

	for _, repo := range repositories {
		if repo.Name == id {
			return &repo, nil
		}
	}

	return nil, nil
}

func (c client) RepositoryUpdate(id string, repo Repository) error {
	data, err := jsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(fmt.Sprintf("%s/%s/%s/%s", repositoryAPIEndpoint, repo.Format, repo.Type, id), data)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) RepositoryDelete(id string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", repositoryAPIEndpoint, id))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}
