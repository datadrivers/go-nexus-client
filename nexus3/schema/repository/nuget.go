package repository

const (
	NugetVersion2 NugetVersion = "V2"
	NugetVersion3 NugetVersion = "V3"
)

type NugetVersion string

// NugetProxy contains data specific to proxy repositories of format Nuget
type NugetProxy struct {
	QueryCacheItemMaxAge int `json:"queryCacheItemMaxAge"`

	// NugetVersion is the used Nuget protocol version
	// Possible values: "V3" or "V2"
	NugetVersion NugetVersion `json:"nugetVersion"`
}
