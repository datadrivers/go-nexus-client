package helm

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	helmHostedAPIEndpoint = helmAPIEndpoint + "/hosted"
)

type RepositoryHelmHostedService struct {
	client *client.Client
}

func NewRepositoryHelmHostedService(c *client.Client) *RepositoryHelmHostedService {
	return &RepositoryHelmHostedService{
		client: c,
	}
}

func (s *RepositoryHelmHostedService) Create(repo repository.HelmHostedRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Post(helmHostedAPIEndpoint, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryHelmHostedService) Get(id string) (*repository.HelmHostedRepository, error) {
	var repo repository.HelmHostedRepository
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", helmHostedAPIEndpoint, id), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	if err := json.Unmarshal(body, &repo); err != nil {
		return nil, fmt.Errorf("could not unmarshal repository: %v", err)
	}
	return &repo, nil
}

func (s *RepositoryHelmHostedService) Update(id string, repo repository.HelmHostedRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", helmHostedAPIEndpoint, id), data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryHelmHostedService) Delete(id string) error {
	return common.DeleteRepository(s.client, id)
}
