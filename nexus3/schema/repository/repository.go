package repository

// Cleanup ...
type Cleanup struct {
	//  Components that match any of the applied policies will be deleted
	PolicyNames []string `json:"policyNames"`
}

// Group contains repository group configuration data
type Group struct {
	MemberNames []string `json:"memberNames,omitempty"`
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
