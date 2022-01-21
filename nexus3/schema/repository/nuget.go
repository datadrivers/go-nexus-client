package repository

const (
	NugetVersion2 NugetVersion = "V2"
	NugetVersion3 NugetVersion = "V3"
)

type NugetVersion string

type NugetGroupRepository struct {
	Name    string `json:"name"`
	Online  bool   `json:"online"`
	Group   `json:"group"`
	Storage `json:"storage"`
}

type NugetHostedRepository struct {
	Name    string        `json:"name"`
	Online  bool          `json:"online"`
	Storage HostedStorage `json:"storage"`

	*Cleanup   `json:"cleanup,omitempty"`
	*Component `json:"component,omitempty"`
}

type NugetProxyRepository struct {
	Name          string `json:"name"`
	Online        bool   `json:"online"`
	Storage       `json:"storage"`
	Proxy         `json:"proxy"`
	NegativeCache `json:"negativeCache"`
	HTTPClient    `json:"httpClient"`
	NugetProxy    `json:"nugetProxy"`

	// RoutingRule is used in POST Call and GET call returns RoutingRuleName. see issue: https://issues.sonatype.org/browse/NEXUS-30973

	// The name of the routing rule assigned to this repository
	RoutingRule *string `json:"routingRule,omitempty"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName *string `json:"routingRuleName,omitempty"`

	*Cleanup `json:"cleanup,omitempty"`
}

// NugetProxy contains data specific to proxy repositories of format Nuget
type NugetProxy struct {
	// How long to cache query results from the proxied repository (in seconds)
	QueryCacheItemMaxAge int `json:"queryCacheItemMaxAge"`
	// NugetVersion is the used Nuget protocol version
	// Possible values: "V3" or "V2"
	NugetVersion NugetVersion `json:"nugetVersion"`
}
