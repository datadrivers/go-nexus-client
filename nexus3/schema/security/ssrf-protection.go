package security

// SsrfProtectionConfiguration represents the Server Side Request Forgery
// protection settings of the Nexus instance.
type SsrfProtectionConfiguration struct {
	// Whether SSRF protection is enabled
	Enabled bool `json:"enabled"`

	// List of IP addresses allowed to bypass SSRF protection
	AllowedIPs []string `json:"allowedIPs,omitempty"`

	// List of domain names allowed to bypass SSRF protection
	AllowedDomains []string `json:"allowedDomains,omitempty"`
}
