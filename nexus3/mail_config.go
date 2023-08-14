package nexus3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema"
)

const (
	MailConfigsAPIEndpoint = basePath + "v1/email"
)

type MailConfigService client.Service

func NewMailConfigService(c *client.Client) *MailConfigService {

	s := &MailConfigService{
		Client: c,
	}
	return s
}

func (s *MailConfigService) Get() (*schema.MailConfig, error) {
	body, resp, err := s.Client.Get(MailConfigsAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}
	var mailconfig schema.MailConfig
	if err := json.Unmarshal(body, &mailconfig); err != nil {
		return nil, fmt.Errorf("could not unmarschal MailConfig: %v", err)
	}
	return &mailconfig, nil
}

func (s *MailConfigService) Create(mailconfig *schema.MailConfig) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(mailconfig)
	if err != nil {
		return err
	}
	body, resp, err := s.Client.Put(MailConfigsAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *MailConfigService) Update(mailconfig *schema.MailConfig) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(mailconfig)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(MailConfigsAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *MailConfigService) Delete() error {
	body, resp, err := s.Client.Delete(MailConfigsAPIEndpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}
	return err
}
