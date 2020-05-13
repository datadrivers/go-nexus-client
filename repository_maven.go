package client

// RepositoryMaven contains additional data of maven repository
type RepositoryMaven struct {
	VersionPolicy string `json:"versionPolicy"`
	LayoutPolicy  string `json:"layoutPolicy"`
}
