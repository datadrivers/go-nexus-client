package common

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/tools"
)

type RepositoryService[R any] struct {
	endpoint string
	client   *client.Client
}

func NewRepositoryService[R any](ep string, c *client.Client) *RepositoryService[R] {
	return &RepositoryService[R]{
		endpoint: ep,
		client:   c,
	}
}

func (s *RepositoryService[R]) Create(repo R) error {
	return s.CreateContext(context.Background(), repo)
}

func (s *RepositoryService[R]) CreateContext(ctx context.Context, repo R) error {

	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.PostContext(ctx, s.endpoint, data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create repository %v: HTTP: %d, %s", repo, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryService[R]) Get(id string) (*R, error) {
	return s.GetContext(context.Background(), id)
}

func (s *RepositoryService[R]) GetContext(ctx context.Context, id string) (*R, error) {
	repo := new(R)
	body, resp, err := s.client.GetContext(ctx, fmt.Sprintf("%s/%s", s.endpoint, id), nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not read repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	if err = json.Unmarshal(body, repo); err != nil {
		return nil, fmt.Errorf("could not unmarshal repository: %v", err)
	}
	return repo, nil
}

func (s *RepositoryService[R]) Update(id string, repo R) error {
	return s.UpdateContext(context.Background(), id, repo)
}

func (s *RepositoryService[R]) UpdateContext(ctx context.Context, id string, repo R) error {
	data, err := tools.JsonMarshalInterfaceToIOReader(repo)
	if err != nil {
		return err
	}
	body, resp, err := s.client.PutContext(ctx, fmt.Sprintf("%s/%s", s.endpoint, id), data)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not update repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}

func (s *RepositoryService[R]) Delete(id string) error {
	return s.DeleteContext(context.Background(), id)
}

func (s *RepositoryService[R]) DeleteContext(ctx context.Context, id string) error {
	return DeleteRepositoryContext(ctx, s.client, id)
}
