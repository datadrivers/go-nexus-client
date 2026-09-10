package license

// LicenseDetails represents the details of the license installed on the Nexus
// instance.
type LicenseDetails struct {
	// The email address of the license contact
	ContactEmail string `json:"contactEmail,omitempty"`

	// The company of the license contact
	ContactCompany string `json:"contactCompany,omitempty"`

	// The name of the license contact
	ContactName string `json:"contactName,omitempty"`

	// The date the license became effective
	EffectiveDate string `json:"effectiveDate,omitempty"`

	// The date the license expires
	ExpirationDate string `json:"expirationDate,omitempty"`

	// The type of the license
	LicenseType string `json:"licenseType,omitempty"`

	// The number of users the license covers
	LicensedUsers string `json:"licensedUsers,omitempty"`

	// The fingerprint of the installed license
	Fingerprint string `json:"fingerprint,omitempty"`

	// The features the license unlocks
	Features string `json:"features,omitempty"`

	// The maximum number of repository components allowed by the license
	MaxRepoComponents *int64 `json:"maxRepoComponents,omitempty"`

	// The maximum number of repository requests allowed by the license
	MaxRepoRequests *int64 `json:"maxRepoRequests,omitempty"`
}
