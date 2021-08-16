package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	realmAPIEndpoint = basePath + "v1/security/realms"
)

// Realm represents an instance of Nexus Realm
type Realm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c client) RealmsActivate(ids []string) error {
	data, err := json.Marshal(ids)
	if err != nil {
		return fmt.Errorf("could not marshal realm IDs: %v", err)
	}

	body, resp, err := c.Put(fmt.Sprintf("%s/active", realmAPIEndpoint), bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("could not activate realms: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not activate realms: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (c client) RealmsActive() ([]string, error) {
	body, resp, err := c.Get(fmt.Sprintf("%s/active", realmAPIEndpoint), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read active realms: HTTP: %d, %v", resp.StatusCode, err)
	}

	var result []string
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("could not unmarshal active realms: %v", err)
	}
	return result, nil
}

func (c client) RealmsAvailable() ([]Realm, error) {
	body, resp, err := c.Get(fmt.Sprintf("%s/available", realmAPIEndpoint), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read available realms: HTTP: %d, %v", resp.StatusCode, err)
	}

	var result []Realm
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("could not unmarshal available realms: %v", err)
	}
	return result, nil
}
