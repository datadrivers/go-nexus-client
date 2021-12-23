package security

// ContentSelector data
type ContentSelector struct {
	// A human-readable description
	Description string `json:"description"`

	// The expression used to identify content
	Expression string `json:"expression"`

	// The content selector name cannot be changed after creation
	Name string `json:"name"`
}
