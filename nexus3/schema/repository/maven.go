package repository

const (
	MavenVersionPolicyRelease  MavenVersionPolicy = "RELEASE"
	MavenVersionPolicySnapshot MavenVersionPolicy = "SNAPSHOT"
	MavenVersionPolicyMixed    MavenVersionPolicy = "MIXED"

	MavenLayoutPolicyStrict     MavenLayoutPolicy = "STRICT"
	MavenLayoutPolicyPermissive MavenLayoutPolicy = "PERMISSIVE"

	MavenContentDispositionInline     MavenContentDisposition = "INLINE"
	MavenContentDispositionAttachment MavenContentDisposition = "ATTACHMENT"
)

// What type of artifacts does this repository store?
type MavenVersionPolicy string

// Validate that all paths are maven artifact or metadata paths
type MavenLayoutPolicy string

// Content Disposition
type MavenContentDisposition string

type MavenGroupRepository struct {
	Name   string `json:"name"`
	Online bool   `json:"online"`
	Maven  *Maven `json:"maven,omitempty"` // Optional because a GET against Nexus API don't return the maven configuration object

	Group   `json:"group"`
	Storage `json:"storage"`
}

type MavenHostedRepository struct {
	Name    string        `json:"name"`
	Online  bool          `json:"online"`
	Storage HostedStorage `json:"storage"`
	Maven   `json:"maven"`

	// Cleanup data
	*Cleanup `json:"cleanup,omitempty"`

	// Components
	*Component `json:"component,omitempty"`
}

type MavenProxyRepository struct {
	Name          string `json:"name"`
	Online        bool   `json:"online"`
	Storage       `json:"storage"`
	Maven         `json:"maven"`
	Proxy         `json:"proxy"`
	NegativeCache `json:"negativeCache"`
	HTTPClient    HTTPClientWithPreemptiveAuth `json:"httpClient"`

	// RoutingRule is used in POST Call and GET call returns RoutingRuleName. see issue: https://issues.sonatype.org/browse/NEXUS-30973

	// The name of the routing rule assigned to this repository
	RoutingRule *string `json:"routingRule,omitempty"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName *string `json:"routingRuleName,omitempty"`

	*Cleanup `json:"cleanup,omitempty"`
}

// Maven contains additional data of maven repository
type Maven struct {
	VersionPolicy      MavenVersionPolicy       `json:"versionPolicy"`
	LayoutPolicy       MavenLayoutPolicy        `json:"layoutPolicy"`
	ContentDisposition *MavenContentDisposition `json:"contentDisposition,omitempty"`
}
