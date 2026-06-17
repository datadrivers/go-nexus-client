// Package capability contains schema types for the Nexus "Capabilities" API.
// See: /service/rest/v1/capabilities
package capability

// Capability represents an existing capability as returned by GET /capabilities
// or after create/update. Each capability has a type and a type-specific set
// of string properties.
type Capability struct {
	ID         string            `json:"id,omitempty"`      // server-assigned
	Type       string            `json:"type"`              // e.g. "baseurl", "httpclient", "outreach", etc.
	Notes      string            `json:"notes,omitempty"`   // free-form
	Enabled    bool              `json:"enabled"`           // on/off
	Properties map[string]string `json:"properties"`        // type-specific key/value pairs
}

// CapabilityCreate is the POST body for creating a capability.
// Same shape as update, but without ID.
type CapabilityCreate struct {
	Type       string            `json:"type"`
	Notes      string            `json:"notes,omitempty"`
	Enabled    bool              `json:"enabled"`
	Properties map[string]string `json:"properties"`
}

// CapabilityUpdate is the PUT body for updating a capability.
// Note: ID must be included in the body for the update to work
type CapabilityUpdate struct {
	ID         string            `json:"id"`              // Required for update
	Type       string            `json:"type"`
	Notes      string            `json:"notes,omitempty"`
	Enabled    bool              `json:"enabled"`
	Properties map[string]string `json:"properties"`
}

// TypeDescriptor (optional) represents an element from GET /capabilities/types
// if you choose to surface type metadata (name/description/etc.). Nexus does
// not guarantee a rich schema here; keep it minimal and tolerant.
type TypeDescriptor struct {
	ID          string `json:"id"`          // capability type id (e.g. "baseurl")
	Name        string `json:"name"`        // human-readable
	Description string `json:"description"` // human-readable
	// Some Nexus builds include per-type property metadata. If you need it later:
	// Properties []PropertyDescriptor `json:"properties,omitempty"`
}

