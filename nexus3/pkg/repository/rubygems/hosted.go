package rubygems

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
	rubyGemsHostedAPIEndpoint = rubyGemsAPIEndpoint + "/hosted"
)

type RepositoryRubyGemsHostedService struct {
	client *client.Client
}

func NewRepositoryRubyGemsHostedService(c *client.Client) *RepositoryRubyGemsHostedService {
	return &RepositoryRubyGemsHostedService{
		client: c,
	}
}

func (s *RepositoryRubyGemsHostedService) Create(repo repository.RubyGemsHostedRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Post(rubyGemsHostedAPIEndpoint, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryRubyGemsHostedService) Get(id string) (*repository.RubyGemsHostedRepository, error) {
	var repo repository.RubyGemsHostedRepository
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", rubyGemsHostedAPIEndpoint, id), nil)
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

func (s *RepositoryRubyGemsHostedService) Update(id string, repo repository.RubyGemsHostedRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", rubyGemsHostedAPIEndpoint, id), data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryRubyGemsHostedService) Delete(id string) error {
	return common.DeleteRepository(s.client, id)
}
