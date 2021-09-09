package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	anonymousAPIEndpoint = basePath + "v1/security/anonymous"
)

// Anonymous config
type AnonymousConfig struct {
	Enabled   bool   `json:"enabled"`
	UserID    string `json:"userId"`
	RealmName string `json:"realmName"`
}

func (c client) AnonymousRead() (*AnonymousConfig, error) {
	body, resp, err := c.Get(anonymousAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read anonymous config: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var anonymous AnonymousConfig
	if err := json.Unmarshal(body, &anonymous); err != nil {
		return nil, fmt.Errorf("could not unmarshal anonymous config: %v", err)
	}

	return &anonymous, nil
}

func (c client) AnonymousUpdate(anonymous AnonymousConfig) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(anonymous)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(anonymousAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not update anonymous config: HTTP %d, %s", resp.StatusCode, string(body))
	}

	return nil
}
