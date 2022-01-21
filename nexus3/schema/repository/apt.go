package repository

type AptHostedRepository struct {
	Name       string        `json:"name"`
	Online     bool          `json:"online"`
	Storage    HostedStorage `json:"storage"`
	Apt        AptHosted     `json:"apt"`
	AptSigning `json:"aptSigning"`

	*Cleanup   `json:"cleanup,omitempty"`
	*Component `json:"component,omitempty"`
}

type AptProxyRepository struct {
	Name          string `json:"name"`
	Online        bool   `json:"online"`
	Storage       `json:"storage"`
	Proxy         `json:"proxy"`
	NegativeCache `json:"negativeCache"`
	HTTPClient    `json:"httpClient"`
	Apt           AptProxy `json:"apt"`

	// RoutingRule is used in POST Call and GET call returns RoutingRuleName. see issue: https://issues.sonatype.org/browse/NEXUS-30973

	// The name of the routing rule assigned to this repository
	RoutingRule *string `json:"routingRule,omitempty"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName *string `json:"routingRuleName,omitempty"`

	*Cleanup `json:"cleanup,omitempty"`
}

// Apt contains data of proxy repositories of format Apt
type AptProxy struct {
	// Distribution to fetch
	Distribution string `json:"distribution"`
	// Whether this repository is flat
	Flat bool `json:"flat"`
}

// Apt contains data of hosted repositories of format Apt
type AptHosted struct {
	// Distribution to fetch
	Distribution string `json:"distribution"`
}

// AptSigning contains signing data of hosted repositores of format Apt
type AptSigning struct {
	// PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)
	Keypair string `json:"keypair"`
	// Passphrase to access PGP signing key
	Passphrase *string `json:"passphrase,omitempty"`
}
