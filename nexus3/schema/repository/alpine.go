package repository

type AlpineGroupRepository struct {
	Name          string `json:"name"`
	Online        bool   `json:"online"`
	Group         `json:"group"`
	Storage       `json:"storage"`
	AlpineSigning `json:"alpineSigning"`
}

type AlpineHostedRepository struct {
	Name    string        `json:"name"`
	Online  bool          `json:"online"`
	Storage HostedStorage `json:"storage"`

	*Cleanup       `json:"cleanup,omitempty"`
	*Component     `json:"component,omitempty"`
	*AlpineSigning `json:"alpineSigning,omitempty"`
}

type AlpineProxyRepository struct {
	Name          string `json:"name"`
	Online        bool   `json:"online"`
	Storage       `json:"storage"`
	Proxy         `json:"proxy"`
	NegativeCache `json:"negativeCache"`
	HTTPClient    `json:"httpClient"`
	AlpineSigning `json:"alpineSigning"`

	// RoutingRule is used in POST Call and GET call returns RoutingRuleName. see issue: https://issues.sonatype.org/browse/NEXUS-30973

	// The name of the routing rule assigned to this repository
	RoutingRule *string `json:"routingRule,omitempty"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName *string `json:"routingRuleName,omitempty"`

	*Cleanup `json:"cleanup,omitempty"`
}

// AlpineSigning contains signing data of repositories of format Alpine
type AlpineSigning struct {
	// RSA signing key pair (PEM armored private key)
	Keypair string `json:"keypair"`
	// Passphrase to access RSA signing key
	Passphrase *string `json:"passphrase,omitempty"`
}
