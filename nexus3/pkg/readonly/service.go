package readonly

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/readonly"
)

const (
	readOnlyAPIEndpoint = client.BasePath + "v1/read-only"
)

var (
	ErrAuthenticationRequired  = errors.New("authentication is required to change readonly state")
	ErrNoChangeToReadOnlyState = errors.New("no change to readonly state")
)

type ReadOnlyService client.Service

func NewReadOnlyService(c *client.Client) *ReadOnlyService {
	return &ReadOnlyService{Client: c}
}

func (s *ReadOnlyService) GetState() (*readonly.State, error) {
	body, resp, err := s.Client.Get(readOnlyAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read readonly state: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var state readonly.State
	if err := json.Unmarshal(body, &state); err != nil {
		return nil, fmt.Errorf("could not unmarshal readonly state : %v", err)
	}

	return &state, nil
}

func (s *ReadOnlyService) Freeze() error {
	_, resp, err := s.Client.Post(fmt.Sprintf("%s/freeze", readOnlyAPIEndpoint), nil)
	if err != nil {
		return err
	}
	return toErr(resp)
}

func (s *ReadOnlyService) Release() error {
	_, resp, err := s.Client.Post(fmt.Sprintf("%s/release", readOnlyAPIEndpoint), nil)
	if err != nil {
		return err
	}
	return toErr(resp)
}

func (s *ReadOnlyService) ForceRelease() error {
	_, resp, err := s.Client.Post(fmt.Sprintf("%s/force-release", readOnlyAPIEndpoint), nil)
	if err != nil {
		return err
	}
	return toErr(resp)
}

func toErr(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusForbidden:
		return ErrAuthenticationRequired
	case http.StatusNotFound:
		return ErrNoChangeToReadOnlyState
	default:
		return errors.New(fmt.Sprintf("unexpected status code %d", resp.StatusCode))
	}
}
