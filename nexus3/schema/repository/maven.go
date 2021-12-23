package repository

// Maven contains additional data of maven repository
type Maven struct {
	VersionPolicy string `json:"versionPolicy"`
	LayoutPolicy  string `json:"layoutPolicy"`
}
