package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

type RepositoryLegacyService struct {
	client *client.Client
}

func NewRepositoryLegacyService(c *client.Client) *RepositoryLegacyService {
	return &RepositoryLegacyService{
		client: c,
	}
}

func jsonUnmarshalRepositories(data []byte) ([]repository.LegacyRepository, error) {
	var repositories []repository.LegacyRepository
	if err := json.Unmarshal(data, &repositories); err != nil {
		return nil, fmt.Errorf("could not unmarshal repositories: %v", err)
	}
	return repositories, nil
}

// Currently only used to replace repository format 'maven2' to 'maven' as API
// returns a format of 'maven2' but requires to send to requests using 'maven'.
func fixRepositoryFormat(s string) string {
	return strings.Replace(s, repository.RepositoryFormatMaven2, "maven", 1)
}

func (s *RepositoryLegacyService) Create(repo repository.LegacyRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Post(fmt.Sprintf("%s/%s/%s", repositoryAPIEndpoint, fixRepositoryFormat(repo.Format), repo.Type), data)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryLegacyService) Get(id string) (*repository.LegacyRepository, error) {
	body, resp, err := s.client.Get(repositoryAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}

	repositories, err := jsonUnmarshalRepositories(body)
	if err != nil {
		return nil, err
	}

	for _, repo := range repositories {
		if repo.Name == id {
			format := repo.Format
			if repo.Format == "maven2" {
				format = "maven"
			}
			body, resp, err := s.client.Get(fmt.Sprintf("%s/%s/%s/%s", repositoryAPIEndpoint, format, repo.Type, repo.Name), nil)
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
	}

	return nil, nil
}

func (s *RepositoryLegacyService) Update(id string, repo repository.LegacyRepository) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}

	body, resp, err := s.client.Put(fmt.Sprintf("%s/%s/%s/%s", repositoryAPIEndpoint, fixRepositoryFormat(repo.Format), repo.Type, id), data)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}

	return nil
}

func (s *RepositoryLegacyService) Delete(id string) error {
	body, resp, err := s.client.Delete(fmt.Sprintf("%s/%s", repositoryAPIEndpoint, id))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}
