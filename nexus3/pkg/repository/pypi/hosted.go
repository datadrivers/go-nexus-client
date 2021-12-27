package pypi

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
	pypiHostedAPIEndpoint = pypiAPIEndpoint + "/hosted"
)

type RepositoryPypiHostedService struct {
	client *client.Client
}

func NewRepositoryPypiHostedService(c *client.Client) *RepositoryPypiHostedService {
	return &RepositoryPypiHostedService{
		client: c,
	}
}

func (s *RepositoryPypiHostedService) Create(repo repository.PypiHostedRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Post(pypiHostedAPIEndpoint, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryPypiHostedService) Get(id string) (*repository.PypiHostedRepository, error) {
	var repo repository.PypiHostedRepository
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", pypiHostedAPIEndpoint, id), nil)
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

func (s *RepositoryPypiHostedService) Update(id string, repo repository.PypiHostedRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", pypiHostedAPIEndpoint, id), data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryPypiHostedService) Delete(id string) error {
	return common.DeleteRepository(s.client, id)
}
