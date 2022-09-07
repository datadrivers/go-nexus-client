package repository

import (
	"encoding/json"
)

const (
	HTTPClientAuthenticationTypeUsername HTTPClientAuthenticationType = "username"
	HTTPClientAuthenticationTypeNtlm     HTTPClientAuthenticationType = "ntlm"
)

// What type of artifacts does this repository store?
type HTTPClientAuthenticationType string

// HTTPClient ...
type HTTPClient struct {
	Authentication *HTTPClientAuthentication `json:"authentication,omitempty"`
	AutoBlock      bool                      `json:"autoBlock"`
	Blocked        bool                      `json:"blocked"`
	Connection     *HTTPClientConnection     `json:"connection,omitempty"`
}

// HTTPClientWithPreemptiveAuth ...
type HTTPClientWithPreemptiveAuth struct {
	// Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive
	AutoBlock bool `json:"autoBlock"`
	// Whether to block outbound connections on the repository
	Blocked bool `json:"blocked"`

	Authentication *HTTPClientAuthenticationWithPreemptive `json:"authentication,omitempty"`
	Connection     *HTTPClientConnection                   `json:"connection,omitempty"`
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

func (hcc *HTTPClientConnection) MarshalJSON() ([]byte, error) {
	type HTTPClientConnectionAlias HTTPClientConnection
	
	// Assign nil if timeout has the default value so JSON marshaler omits it
	if *hcc.Timeout == 0 {
		hcc.Timeout = nil
	}
	
	return json.Marshal(&struct{
		*HTTPClientConnectionAlias
	}{
		HTTPClientConnectionAlias: (*HTTPClientConnectionAlias)(hcc),
	})
}

// HTTPClientAuthentication ...
type HTTPClientAuthentication struct {
	NTLMDomain string                       `json:"ntlmDomain,omitempty"`
	NTLMHost   string                       `json:"ntlmHost,omitempty"`
	Password   string                       `json:"password,omitempty"`
	Type       HTTPClientAuthenticationType `json:"type"`
	Username   string                       `json:"username,omitempty"`
}

// HTTPClientAuthenticationWithPreemptive ...
type HTTPClientAuthenticationWithPreemptive struct {
	NTLMDomain string                      `json:"ntlmDomain,omitempty"`
	NTLMHost   string                      `json:"ntlmHost,omitempty"`
	Password   string                      `json:"password,omitempty"`
	Type       HTTPClientAuthenticationType `json:"type"`
	Username   string                      `json:"username,omitempty"`
	// Whether to use pre-emptive authentication. Use with caution. Defaults to false.
	Preemptive *bool `json:"preemptive,omitempty"`
}
