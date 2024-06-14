package repository

type NpmGroupRepository struct {
	Name    string      `json:"name"`
	Online  bool        `json:"online"`
	Group   GroupDeploy `json:"group"`
	Storage `json:"storage"`
}

type NpmHostedRepository struct {
	Name    string        `json:"name"`
	Online  bool          `json:"online"`
	Storage HostedStorage `json:"storage"`

	*Cleanup   `json:"cleanup,omitempty"`
	*Component `json:"component,omitempty"`
}

type NpmProxyRepository struct {
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

	*Cleanup `json:"cleanup,omitempty"`
	*Npm     `json:"npm,omitempty"`
}

type Npm struct {
	// Remove Non-Cataloged Versions, removed since nexus 3.66.0
	RemoveNonCataloged bool `json:"removeNonCataloged"`
	// Remove Quarantined Versions
	RemoveQuarantined bool `json:"removeQuarantined"`
}
