package client

const (
	NugetVersion2 NugetVersion = "V2"
	NugetVersion3 NugetVersion = "V3"
)

type NugetVersion string

// RepositoryNugetProxy contains data specific to proxy repositories of format Nuget
type RepositoryNugetProxy struct {
	QueryCacheItemMaxAge int `json:"queryCacheItemMaxAge"`

	// NugetVersion is the used Nuget protocol version
	// Possible values: "V3" or "V2"
	NugetVersion NugetVersion `json:"nugetVersion"`
}
