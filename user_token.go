package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	userTokensAPIEndpoint = basePath + "v1/security/user-tokens"
)

// UserTokenConfiguration data structure
type UserTokenConfiguration struct {
	Enabled        bool `json:"enabled"`
	ProtectContent bool `json:"protectContent"`
}

func (c *client) UserTokensApply(userTokens UserTokenConfiguration) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(userTokens)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(userTokensAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if !(resp.StatusCode == http.StatusOK) {
		return fmt.Errorf("could not create/update UserTokenConfiguration configuration: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	return nil
}

func (c *client) UserTokensRead() (*UserTokenConfiguration, error) {
	body, resp, err := c.Get(userTokensAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get UserTokenConfiguration configuration: HTTP: %d, %v", resp.StatusCode, string(body))
	}

	userTokensServer := &UserTokenConfiguration{}
	if err := json.Unmarshal(body, userTokensServer); err != nil {
		return nil, fmt.Errorf("could not unmarshal UserTokenConfiguration configuration: %v", err)
	}

	return userTokensServer, nil
}
