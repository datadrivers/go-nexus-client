package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/schema"
)

const (
	routingRulesAPIEndpoint = basePath + "v1/routing-rules"
)

type RoutingRuleService service

func NewRoutingRuleService(c *client) *RoutingRuleService {

	s := &RoutingRuleService{
		client: c,
	}
	return s
}
func (s *RoutingRuleService) Lists() ([]schema.RoutingRule, error) {
	body, resp, err := s.client.Get(routingRulesAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}

	var rules []schema.RoutingRule
	if err := json.Unmarshal(body, &rules); err != nil {
		return nil, fmt.Errorf("could not unmarschal RoutingRules: %v", err)
	}
	return rules, nil
}

func (s *RoutingRuleService) Get(name string) (*schema.RoutingRule, error) {
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", routingRulesAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}
	var rule schema.RoutingRule
	if err := json.Unmarshal(body, &rule); err != nil {
		return nil, fmt.Errorf("could not unmarschal RoutingRules: %v", err)
	}
	return &rule, nil
}

func (s *RoutingRuleService) Create(rule *schema.RoutingRule) error {
	if err := rule.Mode.IsValid(); err != nil {
		return err
	}
	ioReader, err := jsonMarshalInterfaceToIOReader(rule)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Post(routingRulesAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *RoutingRuleService) Update(rule *schema.RoutingRule) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(rule)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", routingRulesAPIEndpoint, rule.Name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *RoutingRuleService) Delete(name string) error {
	body, resp, err := s.client.Delete(fmt.Sprintf("%s/%s", routingRulesAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}
	return err
}
