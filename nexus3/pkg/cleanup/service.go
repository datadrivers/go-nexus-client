package cleanup

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/cleanuppolicies"
)

const (
	cleanupAPIEndpoint = client.BasePath + "v1/cleanup-policies"
)

type CleanupPolicyService client.Service

func NewCleanupPolicyService(c *client.Client) *CleanupPolicyService {
	return &CleanupPolicyService{Client: c}
}

func (s *CleanupPolicyService) Create(policy *cleanuppolicies.CleanupPolicy) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(policy)
	if err != nil {
		return err
	}
	body, resp, err := s.Client.Post(cleanupAPIEndpoint, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create cleanup policy '%s': HTTP: %d, %s", policy.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *CleanupPolicyService) Get(name string) (*cleanuppolicies.CleanupPolicy, error) {

	body, resp, err := s.Client.Get(fmt.Sprintf("%s/%s", cleanupAPIEndpoint, name), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get cleanup policy '%s': HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	policy := &cleanuppolicies.CleanupPolicy{}
	if err := json.Unmarshal(body, policy); err != nil {
		return nil, fmt.Errorf("could not unmarshal repository: %v", err)
	}

	return policy, nil
}

func (s *CleanupPolicyService) Update(policy *cleanuppolicies.CleanupPolicy) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(policy)
	if err != nil {
		return err
	}
	body, resp, err := s.Client.Put(fmt.Sprintf("%s/%s", cleanupAPIEndpoint, policy.Name), data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update cleanup policy '%s': HTTP: %d, %s", policy.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *CleanupPolicyService) Delete(name string) error {
	endpoint := cleanupAPIEndpoint + "/" + name
	body, resp, err := s.Client.Delete(endpoint)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete cleanup policy '%s': HTTP: %d, %s", name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *CleanupPolicyService) List() ([]*cleanuppolicies.CleanupPolicy, error) {
	body, resp, err := s.Client.Get(cleanupAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not list cleanup policies: HTTP: %d, %s", resp.StatusCode, string(body))
	}
	var policies []*cleanuppolicies.CleanupPolicy
	if err := json.Unmarshal(body, &policies); err != nil {
		return nil, fmt.Errorf("could not unmarshal cleanup policies: %v", err)
	}
	return policies, nil
}
