package repository

// NugetProxy contains data specific to proxy repositories of format Nuget
type NugetProxy struct {
	QueryCacheItemMaxAge int `json:"queryCacheItemMaxAge"`
}
