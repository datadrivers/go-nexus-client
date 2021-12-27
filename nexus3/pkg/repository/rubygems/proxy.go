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
	rubyGemsProxyAPIEndpoint = rubyGemsAPIEndpoint + "/proxy"
)

type RepositoryRubyGemsProxyService struct {
	client *client.Client
}

func NewRepositoryRubyGemsProxyService(c *client.Client) *RepositoryRubyGemsProxyService {
	return &RepositoryRubyGemsProxyService{
		client: c,
	}
}

func (s *RepositoryRubyGemsProxyService) Create(repo repository.RubyGemsProxyRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Post(rubyGemsProxyAPIEndpoint, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryRubyGemsProxyService) Get(id string) (*repository.RubyGemsProxyRepository, error) {
	var repo repository.RubyGemsProxyRepository
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", rubyGemsProxyAPIEndpoint, id), nil)
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

func (s *RepositoryRubyGemsProxyService) Update(id string, repo repository.RubyGemsProxyRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", rubyGemsProxyAPIEndpoint, id), data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryRubyGemsProxyService) Delete(id string) error {
	return common.DeleteRepository(s.client, id)
}
