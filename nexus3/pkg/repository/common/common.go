package common

import (
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
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
