package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	componentAPIEndpoint = "service/rest/v1/components"
)

type ComponentResponse struct {
	Items             []Component `json:"items,omitempty"`
	ContinuationToken interface{} `json:"continuationToken,omitempty"`
}

// Component is the base structure for Nexus Component
type Component struct {
	ID         string `json:"id,omitempty"`
	Repository string `json:"repository,omitempty"`
	Format     string `json:"format,omitempty"`
	Group      string `json:"group,omitempty"`
	Name       string `json:"name,omitempty"`
	Version    string `json:"version,omitempty"`
}

func jsonUnmarshalComponentResponse(data []byte) (*ComponentResponse, error) {
	var componentResponse ComponentResponse
	if err := json.Unmarshal(data, &componentResponse); err != nil {
		return nil, fmt.Errorf("could not unmarshal componentResponse: %v", err)
	}
	return &componentResponse, nil
}

func jsonUnmarshalComponent(data []byte) (*Component, error) {
	var component Component
	if err := json.Unmarshal(data, &component); err != nil {
		return nil, fmt.Errorf("could not unmarshal component: %v", err)
	}
	return &component, nil
}

func (c client) ComponentRead(id string) (*Component, error) {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", componentAPIEndpoint, id))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("could not delete component '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}

	component, err := jsonUnmarshalComponent(body)
	if err != nil {
		return nil, err
	}

	return component, nil
}

func (c client) ComponentUpload(s string, component Component) error {
	panic("implement me")
}

func (c client) ComponentDelete(id string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", componentAPIEndpoint, id))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete component '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (c client) ComponentList(repository string) ([]Component, error) {
	body, resp, err := c.Get(fmt.Sprintf("%s?repository=%s", componentAPIEndpoint, repository), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read repository '%s' component list : HTTP: %d, %s",
			repository, resp.StatusCode, string(body))
	}

	componentResponse, err := jsonUnmarshalComponentResponse(body)
	if err != nil {
		return nil, err
	}

	return componentResponse.Items, nil
}
