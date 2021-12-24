package repository

type GoGroupRepository struct {
	Name    string `json:"name"`
	Online  bool   `json:"online"`
	Group   `json:"group"`
	Storage `json:"storage"`
}

type GoProxyRepository struct {
	Name          string `json:"name"`
	Online        bool   `json:"online"`
	Storage       `json:"storage"`
	Proxy         `json:"proxy"`
	NegativeCache `json:"negativeCache"`
	HTTPClient    `json:"httpClient"`

	RoutingRule *string `json:"routingRule,omitempty"`
	*Cleanup    `json:"cleanup,omitempty"`
	*Npm        `json:"npm,omitempty"`
}
