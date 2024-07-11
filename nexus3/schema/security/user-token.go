package security

// UserTokenConfiguration data structure
type UserTokenConfiguration struct {
	Enabled           bool `json:"enabled"`
	ProtectContent    bool `json:"protectContent"`
	ExpirationEnabled bool `json:"expirationEnabled"`
	ExpirationDays    int  `json:"expirationDays"`
}
