package common

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

const (
	RepositoryAPIEndpoint = client.BasePath + "v1/repositories"
)

func DeleteRepository(client *client.Client, id string) error {
	body, resp, err := client.Delete(fmt.Sprintf("%s/%s", RepositoryAPIEndpoint, id))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func ListRepositories(client *client.Client) ([]repository.RepositoryInfo, error) {
	body, resp, err := client.Get(RepositoryAPIEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not list repository infos: HTTP: %d, %s", resp.StatusCode, string(body))
	}

	var repositoryInfos []repository.RepositoryInfo
	if err := json.Unmarshal(body, &repositoryInfos); err != nil {
		return nil, fmt.Errorf("could not unmarshal list of repository infos: %v", err)
	}
	return repositoryInfos, nil
}
