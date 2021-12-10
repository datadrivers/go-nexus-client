package client

// RepositoryNugetProxy contains data specific to proxy repositories of format Nuget
type RepositoryNugetProxy struct {
	QueryCacheItemMaxAge int `json:"queryCacheItemMaxAge"`

	// NugetVersion is the used Nuget protocol version
	NugetVersion string `json:"nugetVersion"`
}
