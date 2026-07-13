package repository

type AlpineProxyRepository struct {
	Name          string `json:"name"`
	Online        bool   `json:"online"`
	Storage       `json:"storage"`
	Proxy         `json:"proxy"`
	NegativeCache `json:"negativeCache"`
	HTTPClient    `json:"httpClient"`
	AlpineSigning `json:"alpineSigning,omitempty"`

	// RoutingRule is used in POST Call and GET call returns RoutingRuleName. see issue: https://issues.sonatype.org/browse/NEXUS-30973

	// The name of the routing rule assigned to this repository
	RoutingRule *string `json:"routingRule,omitempty"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName *string `json:"routingRuleName,omitempty"`

	*Cleanup `json:"cleanup,omitempty"`
}

// AlpineSigning contains signing data for an Alpine repository
type AlpineSigning struct {
	// PEM-encoded RSA private key used to sign the APKINDEX
	Keypair string `json:"keypair,omitempty"`
	// Passphrase to access the signing key
	Passphrase *string `json:"passphrase,omitempty"`
}
