package schema

import "fmt"

const (
	RoutingRuleModeAllow RoutingRuleMode = "ALLOW"
	RoutingRuleModeBlock RoutingRuleMode = "BLOCK"
)

type RoutingRuleMode string

// RoutingRule is like a filter you can apply to groups in terms of security access and general component retrieval, and can reduce the number of repositories within a group accessed in order to retrieve an component.
type RoutingRule struct {
	// Name of the routing rule
	Name string `json:"name"`

	//Description of the routing rule
	Description string `json:"description,omitempty"`

	// The mode describe how to hande with mathing requests
	// Possible values: "BLOCK" or "ALLOW"
	Mode RoutingRuleMode `json:"mode"`

	// Regular expressions used to identify request paths that are allowed or blocked (depending on above mode)
	Matchers []string `json:"matchers"`
}

// IsValid checks the values of the enum RoutingRuleMode
func (m RoutingRuleMode) IsValid() error {
	switch m {
	case RoutingRuleModeAllow, RoutingRuleModeBlock:
		return nil
	}
	return fmt.Errorf("invalid routing rule mode. Possible values are ALLOW or BLOCK")
}
