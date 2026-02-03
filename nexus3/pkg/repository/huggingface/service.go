package huggingface

import (
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/client"
	"github.com/williamt1997/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/williamt1997/go-nexus-client/nexus3/schema/repository"
)

const (
	huggingfaceAPIEndpoint = common.RepositoryAPIEndpoint + "/huggingface"
)

type (
	RepositoryHuggingfaceProxyService = common.RepositoryService[repository.HuggingfaceProxyRepository]
)

type RepositoryHuggingfaceService struct {
	client *client.Client

	Proxy *RepositoryHuggingfaceProxyService
}

func NewRepositoryHuggingfaceService(c *client.Client) *RepositoryHuggingfaceService {
	return &RepositoryHuggingfaceService{
		client: c,

		Proxy: NewRepositoryHuggingfaceProxyService(c),
	}
}

func NewRepositoryHuggingfaceProxyService(c *client.Client) *RepositoryHuggingfaceProxyService {
	return common.NewRepositoryService[repository.HuggingfaceProxyRepository](huggingfaceAPIEndpoint+"/proxy", c)
}
