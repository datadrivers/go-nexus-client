package security

// SAML data structure
type SAML struct {
	EntityId                   string `json:"entityId,omitempty"`
	IdpMetadata                string `json:"idpMetadata"`
	UsernameAttribute          string `json:"usernameAttribute"`
	FirstNameAttribute         string `json:"firstNameAttribute,omitempty"`
	LastNameAttribute          string `json:"lastNameAttribute,omitempty"`
	EmailAttribute             string `json:"emailAttribute,omitempty"`
	GroupsAttribute            string `json:"groupsAttribute,omitempty"`
	ValidateResponseSignature  bool   `json:"validateResponseSignature,omitempty"`
	ValidateAssertionSignature bool   `json:"validateAssertionSignature,omitempty"`
}
