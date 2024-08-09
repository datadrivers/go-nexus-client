package repository

const (
	DockerProxyIndexTypeHub      DockerProxyIndexType = "HUB"
	DockerProxyIndexTypeRegistry DockerProxyIndexType = "REGISTRY"
	DockerProxyIndexTypeCustom   DockerProxyIndexType = "CUSTOM"
)

type DockerProxyIndexType string

type DockerGroupRepository struct {
	Name    string      `json:"name"`
	Online  bool        `json:"online"`
	Group   GroupDeploy `json:"group"`
	Storage `json:"storage"`
	Docker  `json:"docker"`
}

type DockerHostedRepository struct {
	Name    string              `json:"name"`
	Online  bool                `json:"online"`
	Storage DockerHostedStorage `json:"storage"`
	Docker  `json:"docker"`

	*Cleanup   `json:"cleanup,omitempty"`
	*Component `json:"component,omitempty"`
}

type DockerProxyRepository struct {
	Name          string `json:"name"`
	Online        bool   `json:"online"`
	Storage       `json:"storage"`
	Proxy         `json:"proxy"`
	NegativeCache `json:"negativeCache"`
	HTTPClient    `json:"httpClient"`
	Docker        `json:"docker"`
	DockerProxy   `json:"dockerProxy"`

	// RoutingRule is used in POST Call and GET call returns RoutingRuleName. see issue: https://issues.sonatype.org/browse/NEXUS-30973

	// The name of the routing rule assigned to this repository
	RoutingRule *string `json:"routingRule,omitempty"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName *string `json:"routingRuleName,omitempty"`

	*Cleanup `json:"cleanup,omitempty"`
}

// Docker contains data of a Docker Repositoriy
type Docker struct {
	// Whether to force authentication (Docker Bearer Token Realm required if false)
	ForceBasicAuth bool `json:"forceBasicAuth"`
	// Create an HTTP connector at specified port
	HTTPPort *int `json:"httpPort,omitempty"`
	// Create an HTTPS connector at specified port
	HTTPSPort *int `json:"httpsPort,omitempty"`
	// Whether to allow clients to use the V1 API to interact with this repository
	V1Enabled bool `json:"v1Enabled"`
	// Whether to allow subdomain connector at the specified subdomain
	Subdomain *string `json:"subdomain,omitempty"`
}

// DockerProxy contains data of a Docker Proxy Repository
type DockerProxy struct {
	// Type of Docker Index
	IndexType DockerProxyIndexType `json:"indexType"`
	// Url of Docker Index to use
	IndexURL *string `json:"indexUrl,omitempty"`
	// CacheForeignLayers: Allow Nexus Repository Manager to download and cache foreign layers
	CacheForeignLayers *bool `json:"cacheForeignLayers,omitempty"`
	// ForeignLayerUrlWhitelist is a list of regular expressions used to identify URLs that are allowed for foreign layer requests
	ForeignLayerUrlWhitelist []string `json:"foreignLayerUrlWhitelist"`
}

// DockerHostedStorage contains repository storage for hosted docker
type DockerHostedStorage struct {
	// Blob store used to store repository contents
	BlobStoreName string `json:"blobStoreName"`

	// StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation bool `json:"strictContentTypeValidation"`

	// WritePolicy controls if deployments of and updates to assets are allowed
	WritePolicy StorageWritePolicy `json:"writePolicy"`

	// LatestPolicy: Whether to allow redeploying the 'latest' tag but defer to the Deployment Policy for all other tags
	// WritePolicy ALLOW_ONCE required
	LatestPolicy *bool `json:"latestPolicy,omitempty"`
}
