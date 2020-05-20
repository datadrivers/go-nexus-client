package client

// RepositoryYum contains data of hosted repositories of format Yum
type RepositoryYum struct {
	RepodataDepth int    `json:"repodataDepth"`
	DeployPolicy  string `json:"deployPolicy"`
}
