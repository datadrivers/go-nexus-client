package p2

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
	p2ProxyAPIEndpoint = p2APIEndpoint + "/proxy"
)

type RepositoryP2ProxyService struct {
	client *client.Client
}

func NewRepositoryP2ProxyService(c *client.Client) *RepositoryP2ProxyService {
	return &RepositoryP2ProxyService{
		client: c,
	}
}

func (s *RepositoryP2ProxyService) Create(repo repository.P2ProxyRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Post(p2ProxyAPIEndpoint, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryP2ProxyService) Get(id string) (*repository.P2ProxyRepository, error) {
	var repo repository.P2ProxyRepository
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", p2ProxyAPIEndpoint, id), nil)
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

func (s *RepositoryP2ProxyService) Update(id string, repo repository.P2ProxyRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", p2ProxyAPIEndpoint, id), data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryP2ProxyService) Delete(id string) error {
	return common.DeleteRepository(s.client, id)
}
