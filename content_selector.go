package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	contentSelectorAPIEndpoint = "service/rest/beta/security/content-selectors"
)

// ContentSelector data
type ContentSelector struct {
	Description string `json:"description"`
	Expression  string `json:"expression"`
	Name        string `json:"name"`
}

func (c client) ContentSelectorCreate(cs ContentSelector) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(cs)
	if err != nil {
		return err
	}

	body, resp, err := c.Post(contentSelectorAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create content selector \"%s\": HTTP: %d, %s", cs.Name, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) ContentSelectorRead(name string) (*ContentSelector, error) {
	body, resp, err := c.Get(contentSelectorAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read content selectors: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var contentSelectors []ContentSelector
	if err := json.Unmarshal(body, &contentSelectors); err != nil {
		return nil, fmt.Errorf("could not unmarshal content selectors \"%s\": %v", name, err)
	}

	for _, cs := range contentSelectors {
		if cs.Name == name {
			return &cs, nil
		}
	}

	return nil, nil
}

func (c client) ContentSelectorUpdate(name string, cs ContentSelector) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(cs)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(fmt.Sprintf("%s/%s", contentSelectorAPIEndpoint, name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update content selector \"%s\": HTTP %d, %s", name, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) ContentSelectorDelete(name string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", contentSelectorAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete content selector \"%s\": HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	return nil
}
