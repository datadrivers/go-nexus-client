package cleanuppolicies

type RepositoryFormat string
type CriteriaReleaseType string

const (
	FormatAny       RepositoryFormat = "*"
	FormatApt       RepositoryFormat = "apt"
	FormatBower     RepositoryFormat = "bower"
	FormatCocoapods RepositoryFormat = "cocoapods"
	FormatConan     RepositoryFormat = "conan"
	FormatConda     RepositoryFormat = "conda"
	FormatDocker    RepositoryFormat = "docker"
	FormatGo        RepositoryFormat = "go"
	FormatHelm      RepositoryFormat = "helm"
	FormatMaven     RepositoryFormat = "maven"
	FormatNpm       RepositoryFormat = "npm"
	FormatNuget     RepositoryFormat = "nuget"
	FormatPypi      RepositoryFormat = "pypi"
	FormatR         RepositoryFormat = "r"
	FormatRaw       RepositoryFormat = "raw"
	FormatRubyGems  RepositoryFormat = "rubygems"
	FormatTerraform RepositoryFormat = "terraform"
	FormatYum       RepositoryFormat = "yum"
)

const (
	CriteriaReleaseTypeReleasesAndPrereleases CriteriaReleaseType = "RELEASES_AND_PRERELEASES"
	CriteriaReleaseTypePrereleases            CriteriaReleaseType = "PRERELEASES"
	CriteriaReleaseTypeReleases               CriteriaReleaseType = "RELEASES"
)

type CleanupPolicy struct {
	Notes                   *string              `json:"notes,omitempty"`                   // any details on the specific cleanup policy.
	CriteriaLastBlobUpdated *int                 `json:"criteriaLastBlobUpdated,omitempty"` // the age of the component in days.
	CriteriaLastDownloaded  *int                 `json:"criteriaLastDownloaded,omitempty"`  // the last time the component had been downloaded in days.
	CriteriaReleaseType     *CriteriaReleaseType `json:"criteriaReleaseType,omitempty"`     // is one of: RELEASES_AND_PRERELEASES, PRERELEASES, RELEASES], Only maven2, npm and yum repositories support this field.
	CriteriaAssetRegex      *string              `json:"criteriaAssetRegex,omitempty"`      // a regex string to filter for specific asset paths. Not for gitlfs or *
	Retain                  int                  `json:"retain,omitempty"`                  // number of versions to keep. Only available for Docker and Maven release repositories on PostgreSQL deployments.
	Name                    string               `json:"name"`                              // the name of the policy needs to be unique and cannot be edited once set. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.
	Format                  RepositoryFormat     `json:"format"`                            // one of ["*", "apt", "bower", "cocoapods", "conan", "conda", "docker", "gitlfs", "go", "helm", "maven2", "npm", "nuget", "p2", "pypi", "r", "raw", "rubygems", "yum"]
}
