package repository

const (
	YumDeployPolicyPermissive YumDeployPolicy = "PERMISSIVE"
	YumDeployPolicyStrict     YumDeployPolicy = "STRICT"
)

type YumDeployPolicy string

type YumGroupRepository struct {
	Name    string `json:"name"`
	Online  bool   `json:"online"`
	Group   `json:"group"`
	Storage `json:"storage"`

	*YumSigning `json:"yumSigning,omitempty"`
}

type YumHostedRepository struct {
	Name    string        `json:"name"`
	Online  bool          `json:"online"`
	Storage HostedStorage `json:"storage"`
	Yum     `json:"yum"`

	*Cleanup   `json:"cleanup,omitempty"`
	*Component `json:"component,omitempty"`
}

type YumProxyRepository struct {
	Name          string `json:"name"`
	Online        bool   `json:"online"`
	Storage       `json:"storage"`
	Proxy         `json:"proxy"`
	NegativeCache `json:"negativeCache"`
	HTTPClient    `json:"httpClient"`

	// RoutingRule is used in POST Call and GET call returns RoutingRuleName. see issue: https://issues.sonatype.org/browse/NEXUS-30973

	// The name of the routing rule assigned to this repository
	RoutingRule *string `json:"routingRule,omitempty"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName *string `json:"routingRuleName,omitempty"`

	*Cleanup    `json:"cleanup,omitempty"`
	*YumSigning `json:"yumSigning,omitempty"`
}

// Yum contains data of hosted repositories of format Yum
type Yum struct {
	RepodataDepth int              `json:"repodataDepth"`
	DeployPolicy  *YumDeployPolicy `json:"deployPolicy,omitempty"`
}

type YumSigning struct {
	// PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)
	Keypair *string `json:"keypair,omitempty"`
	// Passphrase to access PGP signing key
	Passphrase *string `json:"passphrase,omitempty"`
}
