package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	repositoryAPIEndpoint = "service/rest/beta/repositories"

	RepositoryFormatApt    = "apt"
	RepositoryFormatBower  = "bower"
	RepositoryFormatConan  = "conan"
	RepositoryFormatDocker = "docker"
	RepositoryFormatGitLFS = "gitlfs"
	RepositoryFormatGo     = "go"
	RepositoryFormatHelm   = "helm"
	RepositoryFormatMaven2 = "maven2"
	RepositoryFormatNPM    = "npm"
	RepositoryFormatNuget  = "nuget"
	RepositoryFormatP2     = "p2"
	RepositoryFormatPyPi   = "pypi"
	RepositoryFormatRAW    = "raw"
	RepositoryFormatRuby   = "rubygems"
	RepositoryFormatYum    = "yum"

	RepositoryTypeGroup  = "group"
	RepositoryTypeHosted = "hosted"
	RepositoryTypeProxy  = "proxy"
)

var (
	// RepositoryFormats contains a list of all supported repository formats
	RepositoryFormats = []string{
		RepositoryFormatApt,
		RepositoryFormatBower,
		RepositoryFormatConan,
		RepositoryFormatDocker,
		RepositoryFormatGitLFS,
		RepositoryFormatGo,
		RepositoryFormatHelm,
		RepositoryFormatMaven2,
		RepositoryFormatNPM,
		RepositoryFormatNuget,
		RepositoryFormatP2,
		RepositoryFormatPyPi,
		RepositoryFormatRAW,
		RepositoryFormatRuby,
		RepositoryFormatYum,
	}

	// RepositoryTypes contains a list of all supported repository types
	RepositoryTypes = []string{
		RepositoryTypeGroup,
		RepositoryTypeHosted,
		RepositoryTypeProxy,
	}
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

	// RepositoryBower data
	*RepositoryBower `json:"bower,omitempty"`

	// RepositoryCleanup data
	*RepositoryCleanup `json:"cleanup,omitempty"`

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

	// Nuget Proxy Repository data
	*RepositoryNugetProxy `json:"nugetProxy,omitempty"`

	// Proxy Repository data
	*RepositoryProxy `json:"proxy,omitempty"`

	// Repository storage data
	*RepositoryStorage `json:"storage,omitempty"`

	// Yum Repository data
	*RepositoryYum `json:"yum,omitempty"`
}

// RepositoryCleanup ...
type RepositoryCleanup struct {
	PolicyNames []string `json:"policyNames"`
}

// RepositoryGroup contains repository group configuration data
type RepositoryGroup struct {
	MemberNames []string `json:"memberNames,omitempty"`
}

// RepositoryHTTPClient ...
type RepositoryHTTPClient struct {
	Authentication *RepositoryHTTPClientAuthentication `json:"authentication,omitempty"`
	AutoBlock      bool                                `json:"autoBlock"`
	Blocked        bool                                `json:"blocked"`
	Connection     *RepositoryHTTPClientConnection     `json:"connection,omitempty"`
}

// RepositoryHTTPClientConnection ...
type RepositoryHTTPClientConnection struct {
	EnableCircularRedirects *bool  `json:"enableCircularRedirects,omitempty"`
	EnableCookies           *bool  `json:"enableCookies,omitempty"`
	Retries                 *int   `json:"retries,omitempty"`
	Timeout                 *int   `json:"timeout,omitempty"`
	UserAgentSuffix         string `json:"userAgentSuffix,omitempty"`
}

// RepositoryHTTPClientAuthentication ...
type RepositoryHTTPClientAuthentication struct {
	NTLMDomain string `json:"ntlmDomain,omitempty"`
	NTLMHost   string `json:"ntlmHost,omitempty"`
	Password   string `json:"password,omitempty"`
	Type       string `json:"type,omitempty"`
	Username   string `json:"username,omitempty"`
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
	RemoteURL      string `json:"remoteUrl,omitempty"`
}

// RepositoryStorage contains repository storage
type RepositoryStorage struct {
	BlobStoreName               string  `json:"blobStoreName,omitempty"`
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

// Currently only used to replace repository format 'maven2' to 'maven' as API
// returns a format of 'maven2' but requires to send to requests using 'maven'.
func fixRepositoryFormat(s string) string {
	return strings.Replace(s, RepositoryFormatMaven2, "maven", 1)
}

func (c client) RepositoryCreate(repo Repository) error {
	data, err := jsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}

	body, resp, err := c.Post(fmt.Sprintf("%s/%s/%s", repositoryAPIEndpoint, fixRepositoryFormat(repo.Format), repo.Type), data)
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

	body, resp, err := c.Put(fmt.Sprintf("%s/%s/%s/%s", repositoryAPIEndpoint, fixRepositoryFormat(repo.Format), repo.Type, id), data)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
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
