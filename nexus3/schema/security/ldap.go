package security

// LDAP data structure
type LDAP struct {
	// The password to bind with. Required if authScheme other than none
	AuthPassword string `json:"authPassword"`

	// The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5
	AuthRealm string `json:"authRealm,omitempty"`

	// Authentication scheme used for connecting to LDAP server
	AuthSchema string `json:"authScheme"`

	// This must be a fully qualified username if simple authentication is used. Required if authScheme other than none
	AuthUserName string `json:"authUsername,omitempty"`

	// How long to wait before retrying
	ConnectionRetryDelaySeconds int32 `json:"connectionRetryDelaySeconds"`

	// How long to wait before timeout
	ConnectionTimeoutSeconds int32 `json:"connectionTimeoutSeconds"`

	// The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN
	GroupBaseDn string `json:"groupBaseDn,omitempty"`

	// This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static
	GroupIDAttribute string `json:"groupIdAttribute,omitempty"`

	// LDAP attribute containing the usernames for the group. Required if groupType is static
	GroupMemberAttribute string `json:"groupMemberAttribute,omitempty"`

	// The format of user ID stored in the group member attribute. Required if groupType is static
	GroupMemberFormat string `json:"groupMemberFormat,omitempty"`

	// LDAP class for group objects. Required if groupType is static
	GroupObjectClass string `json:"groupObjectClass,omitempty"`

	// Are groups located in structures below the group base DN
	GroupSubtree bool `json:"groupSubtree,omitempty"`

	// Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true
	GroupType string `json:"groupType"`

	// LDAP server connection hostname
	Host string `json:"host"`
	ID   string `json:"id"`

	// Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles
	LDAPGroupsAsRoles bool `json:"ldapGroupsAsRoles,omitempty"`

	// How many retry attempts
	MaxIncidentCount int32 `json:"maxIncidentsCount"`

	// LDAP server name
	Name string `json:"name"`

	// LDAP server connection port to use
	Port int32 `json:"port"`

	// LDAP server connection Protocol to use
	Protocol string `json:"protocol"`

	// LDAP location to be added to the connection URL
	SearchBase string `json:"searchBase"`

	// Whether to use certificates stored in Nexus Repository Manager's truststore
	UseTrustStore bool `json:"useTrustStore,omitempty"`

	// The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN
	UserBaseDN string `json:"userBaseDn,omitempty"`

	// This is used to find an email address given the user ID
	UserEmailAddressAttribute string `json:"userEmailAddressAttribute,omitempty"`

	// This is used to find a user given its user ID
	UserIDAttribute string `json:"userIdAttribute,omitempty"`

	// LDAP search filter to limit user search. example: "(|(mail=*@example.com)(uid=dom*))"
	UserLDAPFilter string `json:"userLdapFilter,omitempty"`

	// Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic
	UserMemberOfAttribute string `json:"userMemberOfAttribute,omitempty"`

	// LDAP class for user objects
	UserObjectClass string `json:"userObjectClass,omitempty"`

	// If this field is blank the user will be authenticated against a bind with the LDAP server
	UserPasswordAttribute string `json:"userPasswordAttribute,omitempty"`

	// This is used to find a real name given the user ID
	UserRealNameAttribute string `json:"userRealNameAttribute,omitempty"`

	// Are users located in structures below the user base DN?
	UserSubtree bool `json:"userSubtree,omitempty"`
}
