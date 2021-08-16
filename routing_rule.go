package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	routingRulesAPIEndpoint = basePath + "v1/routing-rules"

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
	return fmt.Errorf("Invalid routing rule mode. Possible values are ALLOW or BLOCK")
}

func (c *client) RoutingRulesLists() ([]RoutingRule, error) {
	body, resp, err := c.Get(routingRulesAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}

	var rules []RoutingRule
	if err := json.Unmarshal(body, &rules); err != nil {
		return nil, fmt.Errorf("could not unmarschal RoutingRules: %v", err)
	}
	return rules, nil
}

func (c *client) RoutingRuleRead(name string) (*RoutingRule, error) {
	body, resp, err := c.Get(fmt.Sprintf("%s/%s", routingRulesAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}
	var rule RoutingRule
	if err := json.Unmarshal(body, &rule); err != nil {
		return nil, fmt.Errorf("could not unmarschal RoutingRules: %v", err)
	}
	return &rule, nil
}

func (c *client) RoutingRuleCreate(rule *RoutingRule) error {
	if err := rule.Mode.IsValid(); err != nil {
		return err
	}
	ioReader, err := jsonMarshalInterfaceToIOReader(rule)
	if err != nil {
		return err
	}
	body, resp, err := c.Post(routingRulesAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (c *client) RoutingRuleUpdate(rule *RoutingRule) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(rule)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(fmt.Sprintf("%s/%s", routingRulesAPIEndpoint, rule.Name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (c *client) RoutingRuleDelete(name string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", routingRulesAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}
	return err
}
