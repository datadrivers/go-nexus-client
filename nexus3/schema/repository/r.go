package repository

type RGroupRepository struct {
	Name    string `json:"name"`
	Online  bool   `json:"online"`
	Group   `json:"group"`
	Storage `json:"storage"`
}

type RHostedRepository struct {
	Name    string        `json:"name"`
	Online  bool          `json:"online"`
	Storage HostedStorage `json:"storage"`

	*Cleanup   `json:"cleanup,omitempty"`
	*Component `json:"component,omitempty"`
}

type RProxyRepository struct {
	Name          string `json:"name"`
	Online        bool   `json:"online"`
	Storage       `json:"storage"`
	Proxy         `json:"proxy"`
	NegativeCache `json:"negativeCache"`
	HTTPClient    `json:"httpClient"`

	RoutingRule *string `json:"routingRule,omitempty"`
	*Cleanup    `json:"cleanup,omitempty"`
}
