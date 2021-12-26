package repository

const (
	DockerProxyIndexTypeHub      DockerProxyIndexType = "HUB"
	DockerProxyIndexTypeRegistry DockerProxyIndexType = "REGISTRY"
	DockerProxyIndexTypeCustom   DockerProxyIndexType = "CUSTOM"
)

type DockerProxyIndexType string

type DockerGroupRepository struct {
	Name    string `json:"name"`
	Online  bool   `json:"online"`
	Group   `json:"group"`
	Storage `json:"storage"`
	Docker  `json:"docker"`
}

type DocekrHostedRepository struct {
	Name    string        `json:"name"`
	Online  bool          `json:"online"`
	Storage HostedStorage `json:"storage"`
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

	RoutingRuleName *string `json:"routingRuleName,omitempty"`
	*Cleanup        `json:"cleanup,omitempty"`
	*Npm            `json:"npm,omitempty"`
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
}

// DockerProxy contains data of a Docker Proxy Repository
type DockerProxy struct {
	// Type of Docker Index
	IndexType *DockerProxyIndexType `json:"indexType"`
	// Url of Docker Index to use
	IndexURL *string `json:"indexUrl,omitempty"`
}