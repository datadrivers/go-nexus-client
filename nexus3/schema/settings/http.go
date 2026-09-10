package settings

// HTTPSettings represents the outbound HTTP system settings of the Nexus
// instance (Administration > System > HTTP).
//
// Optional fields are modelled as pointers so that an unset field is omitted
// from the request payload and left at its Nexus default, rather than being
// sent as a zero value that the API would reject.
type HTTPSettings struct {
	// User-Agent customization appended to the Nexus User-Agent header
	UserAgent *string `json:"userAgent,omitempty"`

	// Connection/Socket timeout in seconds. Between 1 and 3600.
	Timeout *int32 `json:"timeout,omitempty"`

	// Connection/Socket retry attempts. Between 0 and 10.
	Retries *int32 `json:"retries,omitempty"`

	// Outbound HTTP proxy settings
	HTTPProxy *ProxySettings `json:"httpProxy,omitempty"`

	// Outbound HTTPS proxy settings. Nexus only accepts these when the HTTP
	// proxy is enabled as well.
	HTTPSProxy *ProxySettings `json:"httpsProxy,omitempty"`

	// Hosts excluded from proxying
	NonProxyHosts []string `json:"nonProxyHosts,omitempty"`
}

// ProxySettings represents the configuration of a single outbound proxy
// server.
type ProxySettings struct {
	// Whether the proxy is enabled
	Enabled bool `json:"enabled"`

	// The proxy host, without scheme
	Host string `json:"host"`

	// The proxy port. The Nexus API models this as a string.
	Port string `json:"port"`

	// The authentication settings used to reach the proxy
	AuthInfo *AuthSettings `json:"authInfo,omitempty"`
}

// AuthSettings represents the authentication settings used to reach an
// outbound proxy server.
type AuthSettings struct {
	// Whether proxy authentication is enabled
	Enabled bool `json:"enabled"`

	// The username used to authenticate against the proxy
	Username string `json:"username"`

	// The password used to authenticate against the proxy. Nexus never
	// returns the real value on read.
	Password string `json:"password"`

	// The Windows NTLM hostname
	NtlmHost string `json:"ntlmHost,omitempty"`

	// The Windows NTLM domain
	NtlmDomain string `json:"ntlmDomain,omitempty"`
}
