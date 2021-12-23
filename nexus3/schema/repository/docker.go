package repository

// Docker contains data of a Docker Repositoriy
type Docker struct {
	ForceBasicAuth bool `json:"forceBasicAuth"`
	HTTPPort       *int `json:"httpPort,omitempty"`
	HTTPSPort      *int `json:"httpsPort,omitempty"`
	V1Enabled      bool `json:"v1Enabled"`
}

// DockerProxy contains data of a Docker Proxy Repository
type DockerProxy struct {
	IndexType string  `json:"indexType"`
	IndexURL  *string `json:"indexUrl,omitempty"`
}
