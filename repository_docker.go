package client

// RepositoryDocker contains data of a Docker Repositoriy
type RepositoryDocker struct {
	ForceBasicAuth bool `json:"forceBasicAuth"`
	HTTPPort       *int `json:"httpPort"`
	HTTPSPort      *int `json:"httpsPort"`
	V1Enabled      bool `json:"v1Enabled"`
}

// RepositoryDockerProxy contains data of a Docker Proxy Repository
type RepositoryDockerProxy struct {
	IndexType string `json:"indexType"`
	IndexURL  string `json:"indexUrl"`
}

func (c client) RepositoryDockerCreate(repo Repository, repoType string) error {
	return c.RepositoryCreate(repo, "docker", repoType)
}

func (c client) RepositoryDockerRead(id string, repoType string) (*Repository, error) {
	return c.RepositoryRead(id)
}

func (c client) RepositoryDockerUpdate(id string, repo Repository, repoType string) error {
	return c.RepositoryUpdate(id, repo, "docker", repoType)
}

func (c client) RepositoryDockerDelete(id string) error {
	return c.RepositoryDelete(id)
}
