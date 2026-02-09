package repository

// TerraformProxyRepositoryApiRequest represents the request body
// used for creating or updating a Terraform proxy repository.
type TerraformProxyRepositoryApiRequest = TerraformProxyRepository

// TerraformHostedRepositoryApiRequest represents the request body
// used for creating or updating a Terraform hosted repository.
type TerraformHostedRepositoryApiRequest = TerraformHostedRepository

// TerraformProxyRepository represents the configuration of a Terraform proxy repository in Nexus.
type TerraformProxyRepository struct {
	// Repository name
	Name string `json:"name"`

	// Whether the repository is online
	Online bool `json:"online"`

	Storage       `json:"storage"`
	Proxy         `json:"proxy"`
	NegativeCache `json:"negativeCache"`
	HTTPClient    `json:"httpClient"`

	// The name of the routing rule assigned to this repository
	RoutingRule *string `json:"routingRule,omitempty"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName *string `json:"routingRuleName,omitempty"`

	// Cleanup policy configuration (optional)
	*Cleanup `json:"cleanup,omitempty"`
}


// TerraformSigningAttributes represents terraformSigning section
// required by TerraformHostedRepositoryApiRequest.
type TerraformSigningAttributes struct {
	// PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor). Required.
	Keypair string `json:"keypair"`

	// Passphrase to access PGP signing key (optional)
	Passphrase *string `json:"passphrase,omitempty"`
}

// TerraformHostedRepository represents the configuration of a Terraform hosted repository in Nexus.
type TerraformHostedRepository struct {
	// A unique identifier for this repository
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage HostedStorage `json:"storage"`

	// Cleanup policy configuration (optional)
	*Cleanup `json:"cleanup,omitempty"`

	// Component attributes (optional) - used by Nexus Firewall namespace-confusion protection
	Component *Component `json:"component,omitempty"`

	// Swagger: TerraformSigningAttributes (required)
	TerraformSigning TerraformSigningAttributes `json:"terraformSigning"`
}

