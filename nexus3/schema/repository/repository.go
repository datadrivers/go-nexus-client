package repository

// Cleanup ...
type Cleanup struct {
	PolicyNames []string `json:"policyNames"`
}

// Group contains repository group configuration data
type Group struct {
	MemberNames []string `json:"memberNames,omitempty"`
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
	EnableCircularRedirects *bool  `json:"enableCircularRedirects,omitempty"`
	EnableCookies           *bool  `json:"enableCookies,omitempty"`
	Retries                 *int   `json:"retries,omitempty"`
	Timeout                 *int   `json:"timeout,omitempty"`
	UserAgentSuffix         string `json:"userAgentSuffix,omitempty"`
	UseTrustStore           *bool  `json:"useTrustStore,omitempty"`
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
	Enabled bool `json:"enabled"`
	TTL     int  `json:"timeToLive"`
}

// Proxy contains Proxy Repository data
type Proxy struct {
	ContentMaxAge  int    `json:"contentMaxAge"`
	MetadataMaxAge int    `json:"metadataMaxAge"`
	RemoteURL      string `json:"remoteUrl,omitempty"`
}

// Storage contains repository storage
type Storage struct {
	BlobStoreName               string  `json:"blobStoreName,omitempty"`
	StrictContentTypeValidation bool    `json:"strictContentTypeValidation"`
	WritePolicy                 *string `json:"writePolicy,omitempty"`
}

// Component ...
type Component struct {
	ProprietaryComponents bool `json:"proprietaryComponents"`
}
