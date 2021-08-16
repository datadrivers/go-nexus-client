package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	scriptsAPIEndpoint = basePath + "v1/script"
)

// Script describe a groovy script object that can be run on the nexus server
type Script struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

func (c *client) ScriptLists() ([]Script, error) {
	body, resp, err := c.Get(scriptsAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}

	var scripts []Script
	if err := json.Unmarshal(body, &scripts); err != nil {
		return nil, fmt.Errorf("could not unmarschal scripts: %v", err)
	}
	return scripts, nil
}

func (c *client) ScriptRead(name string) (*Script, error) {
	body, resp, err := c.Get(fmt.Sprintf("%s/%s", scriptsAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}
	var script Script
	if err := json.Unmarshal(body, &script); err != nil {
		return nil, fmt.Errorf("could not unmarschal scripts: %v", err)
	}
	return &script, nil
}

func (c *client) ScriptCreate(script *Script) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(script)
	if err != nil {
		return err
	}
	body, resp, err := c.Post(scriptsAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (c *client) ScriptUpdate(script *Script) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(script)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(fmt.Sprintf("%s/%s", scriptsAPIEndpoint, script.Name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (c *client) ScriptDelete(name string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", scriptsAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}
	return err
}

func (c *client) ScriptRun(name string) error {
	body, resp, err := c.Post(fmt.Sprintf("%s/%s/run", scriptsAPIEndpoint, name), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", string(body))
	}
	return err
}
