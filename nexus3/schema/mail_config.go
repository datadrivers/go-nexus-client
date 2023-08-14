package schema

type MailConfig struct {
	// Toogle if mail config is active or not
	Enabled *bool `json:"enabled,omitempty"`

	// Host
	Host string `json:"host"`

	// Port
	Port int `json:"port"`

	// Username
	Username *string `json:"username,omitempty"`

	// Password
	Password *string `json:"password,omitempty"`

	// FromAddress
	FromAddress string `json:"fromAddress"`

	// Subject Prefix
	SubjectPrefix *string `json:"subjectPrefix,omitempty"`

	// StartTlsEnabled
	StartTlsEnabled *bool `json:"startTlsEnabled,omitempty"`

	// StartTlsRequired
	StartTlsRequired *bool `json:"startTlsRequired,omitempty"`

	// sslOnConectEnabled
	SslOnConnectEnabled *bool `json:"sslOnConnectEnabled,omitempty"`

	// sslServerIdentityCheckEnabled
	SslServerIdentityCheckEnabled *bool `json:"sslServerIdentityCheckEnabled,omitempty"`

	// nexusTrustStoreEnabled
	NexusTrustStoreEnabled *bool `json:"nexusTrustStoreEnabled,omitempty"`
}
