package repository

const (
	HTTPClientAuthenticationTypeUsername HTTPClientAuthenticationType = "username"
	HTTPClientAuthenticationTypeNtlm     HTTPClientAuthenticationType = "ntlm"
)

// What type of artifacts does this repository store?
type HTTPClientAuthenticationType string

// Cleanup ...
type Cleanup struct {
	//  Components that match any of the applied policies will be deleted
	PolicyNames []string `json:"policyNames"`
}

// Group contains repository group configuration data
type Group struct {
	MemberNames []string `json:"memberNames,omitempty"`
}

// HTTPClientWithPreemptiveAuth
type HTTPClientWithPreemptiveAuth struct {
	// Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive
	AutoBlock bool `json:"autoBlock"`
	// Whether to block outbound connections on the repository
	Blocked bool `json:"blocked"`

	Authentication *HTTPClientAuthenticationWithPreemptive `json:"authentication,omitempty"`
	Connection     *HTTPClientConnection                   `json:"connection,omitempty"`
}

// HTTPClient ...
type HTTPClient struct {
	Authentication *HTTPClientAuthentication `json:"authentication,omitempty"`
	AutoBlock      bool                      `json:"autoBlock"`
	Blocked        bool                      `json:"blocked"`
	Connection     *HTTPClientConnection     `json:"connection,omitempty"`
}

// HTTPClientConnection ...
type HTTPClientConnection struct {
	// Whether to enable redirects to the same location (may be required by some servers)
	EnableCircularRedirects *bool `json:"enableCircularRedirects,omitempty"`
	// Whether to allow cookies to be stored and used
	EnableCookies *bool `json:"enableCookies,omitempty"`
	// Total retries if the initial connection attempt suffers a timeout
	Retries *int `json:"retries,omitempty"`
	// Seconds to wait for activity before stopping and retrying the connection",
	Timeout *int `json:"timeout,omitempty"`
	// Custom fragment to append to User-Agent header in HTTP requests
	UserAgentSuffix string `json:"userAgentSuffix,omitempty"`
	// Use certificates stored in the Nexus Repository Manager truststore to connect to external systems
	UseTrustStore *bool `json:"useTrustStore,omitempty"`
}

// HTTPClientAuthenticationWithPreemptive ...
type HTTPClientAuthenticationWithPreemptive struct {
	NTLMDomain *string                       `json:"ntlmDomain,omitempty"`
	NTLMHost   *string                       `json:"ntlmHost,omitempty"`
	Password   *string                       `json:"password,omitempty"`
	Type       *HTTPClientAuthenticationType `json:"type,omitempty"`
	Username   *string                       `json:"username,omitempty"`
	// Whether to use pre-emptive authentication. Use with caution. Defaults to false.
	Preemptive *bool `json:"preemptive,omitempty"`
}

// HTTPClientAuthentication ...
type HTTPClientAuthentication struct {
	NTLMDomain string `json:"ntlmDomain,omitempty"`
	NTLMHost   string `json:"ntlmHost,omitempty"`
	Password   string `json:"password,omitempty"`
	Type       string `json:"type,omitempty"`
	Username   string `json:"username,omitempty"`
}

// NegativeCache ...
type NegativeCache struct {
	// Whether to cache responses for content not present in the proxied repository
	Enabled bool `json:"enabled"`

	// How long to cache the fact that a file was not found in the repository (in minutes)
	TTL int `json:"timeToLive"`
}

// Proxy contains Proxy Repository data
type Proxy struct {
	// How long to cache artifacts before rechecking the remote repository (in minutes)
	ContentMaxAge int `json:"contentMaxAge"`

	// How long to cache metadata before rechecking the remote repository (in minutes)
	MetadataMaxAge int `json:"metadataMaxAge"`

	// Location of the remote repository being proxied
	RemoteURL *string `json:"remoteUrl,omitempty"`
}

// Component ...
type Component struct {
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents bool `json:"proprietaryComponents"`
}
