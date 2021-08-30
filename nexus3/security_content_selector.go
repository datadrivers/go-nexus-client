package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	securityContentSelectorAPIEndpoint = securityAPIEndpoint + "/content-selectors"
)

type SecurityContentSelectorService service

func NewSecurityContentSelectorService(c *client) *SecurityContentSelectorService {

	s := &SecurityContentSelectorService{
		client: c,
	}
	return s
}

// SecurityContentSelector data
type SecurityContentSelector struct {
	// A human-readable description
	Description string `json:"description"`

	// The expression used to identify content
	Expression string `json:"expression"`

	// The content selector name cannot be changed after creation
	Name string `json:"name"`
}

func (s SecurityContentSelectorService) Create(cs SecurityContentSelector) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(cs)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(securityContentSelectorAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create content selector \"%s\": HTTP: %d, %s", cs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (s SecurityContentSelectorService) List() ([]SecurityContentSelector, error) {
	body, resp, err := s.client.Get(securityContentSelectorAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read content selectors: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var contentSelectors []SecurityContentSelector
	if err := json.Unmarshal(body, &contentSelectors); err != nil {
		return nil, fmt.Errorf("could not unmarshal content selector list: %v", err)
	}

	return contentSelectors, nil
}

func (s SecurityContentSelectorService) Get(name string) (*SecurityContentSelector, error) {
	contentSelectors, err := s.List()
	if err != nil {
		return nil, err
	}

	for _, cs := range contentSelectors {
		if cs.Name == name {
			return &cs, nil
		}
	}

	return nil, nil
}

func (s SecurityContentSelectorService) Update(name string, cs SecurityContentSelector) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(cs)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", securityContentSelectorAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update content selector \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (s SecurityContentSelectorService) Delete(name string) error {
	body, resp, err := s.client.Delete(fmt.Sprintf("%s/%s", securityContentSelectorAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete content selector \"%s\": HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	return nil
}
