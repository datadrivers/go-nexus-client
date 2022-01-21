package repository

const (
	RawContentDispositionInline     RawContentDisposition = "INLINE"
	RawContentDispositionAttachment RawContentDisposition = "ATTACHMENT"
)

// Content Disposition
type RawContentDisposition string

type RawGroupRepository struct {
	Name   string `json:"name"`
	Online bool   `json:"online"`

	Group   `json:"group"`
	Storage `json:"storage"`

	*Raw `json:"raw,omitempty"`
}

type RawHostedRepository struct {
	Name    string        `json:"name"`
	Online  bool          `json:"online"`
	Storage HostedStorage `json:"storage"`

	*Cleanup   `json:"cleanup,omitempty"`
	*Component `json:"component,omitempty"`
	*Raw       `json:"raw,omitempty"`
}

type RawProxyRepository struct {
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
	*Raw     `json:"raw,omitempty"`
}

type Raw struct {
	ContentDisposition *RawContentDisposition `json:"contentDisposition,omitempty"`
}
