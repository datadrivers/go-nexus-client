package security

// AnonymousAccessSettings config
type AnonymousAccessSettings struct {
	// Whether or not Anonymous Access is enabled
	Enabled bool `json:"enabled"`

	// The username of the anonymous account
	UserID string `json:"userId"`

	// The name of the authentication realm for the anonymous account
	RealmName string `json:"realmName"`
}
