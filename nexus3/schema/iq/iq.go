package iq

// IQServerConfiguration represents the IQ Server connection configuration
type IQServerConfiguration struct {
	Enabled              bool              `json:"enabled"`
	ShowLink             bool              `json:"showLink"`
	URL                  *string           `json:"url"`
	AuthenticationType   *string           `json:"authenticationType"` // "USER" or "PKI"
	Username             *string           `json:"username,omitempty"`
	Password             *string           `json:"password,omitempty"`
	UseTrustStoreForURL  bool              `json:"useTrustStoreForUrl"`
	TimeoutSeconds       *int              `json:"timeoutSeconds"`
	Properties           *string           `json:"properties,omitempty"`
	FailOpenModeEnabled  bool              `json:"failOpenModeEnabled"`
}
