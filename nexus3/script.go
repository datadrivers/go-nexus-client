package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/schema"
)

const (
	scriptsAPIEndpoint = basePath + "v1/script"
)

type ScriptService service

func NewScriptService(c *client) *ScriptService {

	s := &ScriptService{
		client: c,
	}
	return s
}

func (s *ScriptService) List() ([]schema.Script, error) {
	body, resp, err := s.client.Get(scriptsAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}

	var scripts []schema.Script
	if err := json.Unmarshal(body, &scripts); err != nil {
		return nil, fmt.Errorf("could not unmarschal scripts: %v", err)
	}
	return scripts, nil
}

func (s *ScriptService) Get(name string) (*schema.Script, error) {
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", scriptsAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}
	var script schema.Script
	if err := json.Unmarshal(body, &script); err != nil {
		return nil, fmt.Errorf("could not unmarschal scripts: %v", err)
	}
	return &script, nil
}

func (s *ScriptService) Create(script *schema.Script) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(script)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Post(scriptsAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *ScriptService) Update(script *schema.Script) error {
	ioReader, err := jsonMarshalInterfaceToIOReader(script)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", scriptsAPIEndpoint, script.Name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *ScriptService) Delete(name string) error {
	body, resp, err := s.client.Delete(fmt.Sprintf("%s/%s", scriptsAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}
	return err
}

func (s *ScriptService) Run(name string) error {
	body, resp, err := s.client.Post(fmt.Sprintf("%s/%s/run", scriptsAPIEndpoint, name), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", string(body))
	}
	return err
}
