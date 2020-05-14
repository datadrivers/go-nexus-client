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
	IndexType string  `json:"indexType"`
	IndexURL  *string `json:"indexUrl,omitempty"`
}
