package schema

type MailConfig struct {
	// Toogle if mail config is active or not
	Enabled bool `json:"enabled"`

	// Host
	Host string `json:"host"`

	// Port
	Port int `json:"port"`

	// Username
	Username string `json:"username"`

	// Password
	Password string `json:"password"`

	// FromAddress
	FromAddress string `json:"fromAddress"`

	// Subject Prefix
	SubjectPrefix string `json:"subjectPrefix"`

	// StartTlsEnabled
	StartTlsEnabled bool `json:"startTlsEnabled"`

	// StartTlsRequired
	StartTlsRequired bool `json:"startTlsRequired"`

	// sslOnConectEnabled
	SslOnConnectEnabled bool `json:"sslOnConnectEnabled"`

	// sslServerIdentityCheckEnabled
	SslServerIdentityCheckEnabled bool `json:"sslServerIdentityCheckEnabled"`

	// nexusTrustStoreEnabled
	NexusTrustStoreEnabled bool `json:"nexusTrustStoreEnabled"`
}
