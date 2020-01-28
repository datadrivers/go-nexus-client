package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	repositoryAPIEndpoint = "service/rest/beta/repositories"
)

// Repository ...
type Repository struct {
	Cleanup RepositoryCleanup `json:"cleanup"`
	Format  string            `json:"format,omitempty"`
	Name    string            `json:"name"`
	Online  bool              `json:"online"`
	Storage RepositoryStorage `json:"storage"`
	Type    string            `json:"type,omitempty"`

	// Apt Repository data
	*RepositoryApt        `json:"apt,omitempty"`
	*RepositoryAptSigning `json:"aptSigning,omitempty"`

	// Docker Repository data
	*RepositoryDocker `json:"docker"`
}

// RepositoryCleanup ...
type RepositoryCleanup struct {
	PolicyNames []string `json:"policyNames"`
}

// RepositoryStorage ...
type RepositoryStorage struct {
	BlobStoreName               string `json:"blobStoreName"`
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation"`
	WritePolicy                 string `json:"writePolicy"`
}

func (c client) RepositoryCreate(repo Repository, format string, repoType string) error {
	data, err := jsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}

	body, resp, err := c.Post(fmt.Sprintf("%s/%s/%s", repositoryAPIEndpoint, format, repoType), data)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository '%s': HTTP: %d, %s", repo.Name, resp.StatusCode, string(body))
	}
	return nil
}

func (c client) RepositoryRead(id string) (*Repository, error) {
	body, resp, err := c.Get(fmt.Sprintf("%s", repositoryAPIEndpoint), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}

	var repositories []Repository
	if err := json.Unmarshal(body, &repositories); err != nil {
		return nil, fmt.Errorf("could not unmarshal repositories: %v", err)
	}

	for _, repo := range repositories {
		if repo.Name == id {
			return &repo, nil
		}
	}

	return nil, nil
}

func (c client) RepositoryUpdate(id string, repo Repository, format string, repoType string) error {
	data, err := jsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}

	body, resp, err := c.Put(fmt.Sprintf("%s/%s/%s/%s", repositoryAPIEndpoint, format, repoType, id), data)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}

	return nil
}

func (c client) RepositoryDelete(id string) error {
	body, resp, err := c.Delete(fmt.Sprintf("%s/%s", repositoryAPIEndpoint, id))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}
