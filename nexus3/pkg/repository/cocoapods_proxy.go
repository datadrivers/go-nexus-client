package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	cocoapodsProxyAPIEndpoint = repositoryAPIEndpoint + "/cocoapods/proxy"
)

type RepositoryCocoapodsProxyService struct {
	client *client.Client
}

func NewRepositoryCocoapodsProxyService(c *client.Client) *RepositoryCocoapodsProxyService {
	return &RepositoryCocoapodsProxyService{
		client: c,
	}
}

func (s *RepositoryCocoapodsProxyService) Create(repo repository.CocoapodsProxyRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Post(cocoapodsProxyAPIEndpoint, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryCocoapodsProxyService) Get(id string) (*repository.CocoapodsProxyRepository, error) {
	var repo repository.CocoapodsProxyRepository
	body, resp, err := s.client.Get(fmt.Sprintf("%s/%s", cocoapodsProxyAPIEndpoint, id), nil)
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

func (s *RepositoryCocoapodsProxyService) Update(id string, repo repository.CocoapodsProxyRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s", cocoapodsProxyAPIEndpoint, id), data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryCocoapodsProxyService) Delete(id string) error {
	return deleteRepository(s.client, id)
}
