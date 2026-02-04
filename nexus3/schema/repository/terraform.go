package repository

// TerraformProxyRepositoryApiRequest represents the request body
// used for creating or updating a Terraform proxy repository.
type TerraformProxyRepositoryApiRequest = TerraformProxyRepository

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
