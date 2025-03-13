package security

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

const (
	securityUsersAPIEndpoint = securityAPIEndpoint + "/users"
)

type SecurityUserService client.Service

func NewSecurityUserService(c *client.Client) *SecurityUserService {

	s := &SecurityUserService{
		Client: c,
	}
	return s
}

func jsonUnmarshalUsers(data []byte) ([]security.User, error) {
	var users []security.User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, fmt.Errorf("could not unmarschal users: %v", err)
	}
	return users, nil
}

func (s *SecurityUserService) Create(user security.User) error {
	ioReader, err := tools.JsonMarshalInterfaceToIOReader(user)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Post(securityUsersAPIEndpoint, ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *SecurityUserService) Get(id string, source *string) (*security.User, error) {
	baseURL, err := url.Parse(securityUsersAPIEndpoint)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("userId", id)
	if source != nil {
		params.Add("source", *source)
	}

	body, resp, err := s.Client.Get(baseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}

	users, err := jsonUnmarshalUsers(body)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.UserID == id {
			return &user, nil
		}
	}

	return nil, nil
}

func (s *SecurityUserService) Update(id string, user security.User) error {
	// Not sure what this is and why is required to update a user
	if user.Source == "" {
		user.Source = "default"
	}

	ioReader, err := tools.JsonMarshalInterfaceToIOReader(user)
	if err != nil {
		return err
	}

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/%s", securityUsersAPIEndpoint, id), ioReader)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (s *SecurityUserService) Delete(id string) error {
	body, resp, err := s.Client.Delete(fmt.Sprintf("%s/%s", securityUsersAPIEndpoint, id))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("%s", string(body))
	}
	return err
}

func (s *SecurityUserService) ChangePassword(id string, password string) error {
	data := bytes.NewReader([]byte(password))
	// UserChangePassword  must be send with content-type text/plain :-/
	s.Client.ContentTypeTextPlain()
	defer s.Client.ContentTypeJSON()

	body, resp, err := s.Client.Put(fmt.Sprintf("%s/%s/change-password", securityUsersAPIEndpoint, id), data)
	if err != nil {
		return fmt.Errorf("could not change password of user '%s': %v", id, err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not change password of user '%s':  HTTP: %d, %s ", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *SecurityUserService) List(source *string) ([]security.User, error) {
	baseURL, err := url.Parse(securityUsersAPIEndpoint)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	if source != nil {
		params.Add("source", *source)
	}

	baseURL.RawQuery = params.Encode()

	body, resp, err := s.Client.Get(baseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", string(body))
	}

	users, err := jsonUnmarshalUsers(body)
	if err != nil {
		return nil, err
	}

	return users, nil
}
