package common

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/stretchr/testify/assert"
)

type testRepository struct {
	Name string `json:"name"`
}

func newTestService(t *testing.T, status int, body string) *RepositoryService[testRepository] {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)
	return NewRepositoryService[testRepository]("service/rest/v1/repositories/raw/hosted", client.NewClient(client.Config{URL: server.URL}))
}

func TestGetReturnsRepository(t *testing.T) {
	repo, err := newTestService(t, http.StatusOK, `{"name":"example"}`).Get("example")
	assert.NoError(t, err)
	assert.Equal(t, "example", repo.Name)
}

func TestGetReturnsNilForMissingRepository(t *testing.T) {
	repo, err := newTestService(t, http.StatusNotFound, `{"message":"Repository not found"}`).Get("example")
	assert.NoError(t, err)
	assert.Nil(t, repo)
}

func TestGetReturnsErrorForOtherStatus(t *testing.T) {
	repo, err := newTestService(t, http.StatusForbidden, "").Get("example")
	assert.Error(t, err)
	assert.Nil(t, repo)
}
