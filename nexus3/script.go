package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/tools"
	"github.com/williamt1997/go-nexus-client/nexus3/schema"
)

const (
	scriptsAPIEndpoint = basePath + "v1/script"
)

type ScriptService client.Service

func NewScriptService(c *client.Client) *ScriptService {

	s := &ScriptService{
		Client: c,
	}
	return s
}

func (s *ScriptService) List() ([]schema.Script, error) {
	body, resp, err := s.Client.Get(scriptsAPIEndpoint, nil)
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
	body, resp, err := s.Client.Get(fmt.Sprintf("%s/%s", scriptsAPIEndpoint, name), nil)
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
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(script)
	if err != nil {
		return err
	}
	body, resp, err := s.Client.Post(scriptsAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *ScriptService) Update(script *schema.Script) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(script)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/%s", scriptsAPIEndpoint, script.Name), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *ScriptService) Delete(name string) error {
	body, resp, err := s.Client.Delete(fmt.Sprintf("%s/%s", scriptsAPIEndpoint, name))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}
	return err
}

func (s *ScriptService) Run(name string) error {
	body, resp, err := s.Client.Post(fmt.Sprintf("%s/%s/run", scriptsAPIEndpoint, name), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", string(body))
	}
	return err
}
func (s *ScriptService) RunWithPayload(name, payload string) error {
	r := strings.NewReader(payload)
	body, resp, err := s.Client.Post(fmt.Sprintf("%s/%s/run", scriptsAPIEndpoint, name), r)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", string(body))
	}
	return err
}
