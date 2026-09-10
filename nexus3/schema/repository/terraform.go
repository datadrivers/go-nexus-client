package repository

type TerraformGroupRepository struct {
	Name      string `json:"name"`
	Online    bool   `json:"online"`
	Group     `json:"group"`
	Storage   `json:"storage"`
	Terraform `json:"terraform"`
}

type TerraformHostedRepository struct {
	Name             string        `json:"name"`
	Online           bool          `json:"online"`
	Storage          HostedStorage `json:"storage"`
	TerraformSigning `json:"terraformSigning"`

	*Cleanup   `json:"cleanup,omitempty"`
	*Component `json:"component,omitempty"`
}

type TerraformProxyRepository struct {
	Name          string `json:"name"`
	Online        bool   `json:"online"`
	Storage       `json:"storage"`
	Proxy         `json:"proxy"`
	NegativeCache `json:"negativeCache"`
	HTTPClient    HTTPClientWithPreemptiveAuth `json:"httpClient"`
	Terraform     `json:"terraform"`

	// RoutingRule is used in POST Call and GET call returns RoutingRuleName. see issue: https://issues.sonatype.org/browse/NEXUS-30973

	// The name of the routing rule assigned to this repository
	RoutingRule *string `json:"routingRule,omitempty"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName *string `json:"routingRuleName,omitempty"`

	*Cleanup `json:"cleanup,omitempty"`
}

// Terraform contains additional data of terraform repositories
type Terraform struct {
	// Indicates if this repository requires authentication overriding anonymous access.
	RequireAuthentication bool `json:"requireAuthentication"`
}

// TerraformSigning contains signing data of hosted repositories of format Terraform
type TerraformSigning struct {
	// PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)
	Keypair string `json:"keypair"`
	// Passphrase to access PGP signing key
	Passphrase *string `json:"passphrase,omitempty"`
}
