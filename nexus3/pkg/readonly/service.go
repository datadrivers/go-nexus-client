package readonly

import (
	"encoding/json"
	"fmt"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/readonly"
	"net/http"
)

const (
	readOnlyAPIEndpoint = client.BasePath + "v1/read-only"
)

type ReadOnlyService client.Service

func NewReadOnlyService(c *client.Client) *ReadOnlyService {
	return &ReadOnlyService{}
}

func (s *ReadOnlyService) GetState() (*readonly.State, error) {
	body, resp, err := s.Client.Get(readOnlyAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read read only state: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var state readonly.State
	if err := json.Unmarshal(body, &state); err != nil {
		return nil, fmt.Errorf("could not unmarshal read only state : %v", err)
	}

	return &state, nil
}

func (s *ReadOnlyService) Freeze() error {
	return nil
}

func (s *ReadOnlyService) Release() error {
	return nil
}

func (s *ReadOnlyService) ForceRelease() error {
	return nil
}
