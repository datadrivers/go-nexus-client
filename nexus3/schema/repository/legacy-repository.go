package repository

const (
	RepositoryFormatApt    = "apt"
	RepositoryFormatBower  = "bower"
	RepositoryFormatConan  = "conan"
	RepositoryFormatDocker = "docker"
	RepositoryFormatGitLFS = "gitlfs"
	RepositoryFormatGo     = "go"
	RepositoryFormatHelm   = "helm"
	RepositoryFormatMaven2 = "maven2"
	RepositoryFormatNPM    = "npm"
	RepositoryFormatNuget  = "nuget"
	RepositoryFormatP2     = "p2"
	RepositoryFormatPyPi   = "pypi"
	RepositoryFormatRAW    = "raw"
	RepositoryFormatRuby   = "rubygems"
	RepositoryFormatYum    = "yum"

	RepositoryTypeGroup  = "group"
	RepositoryTypeHosted = "hosted"
	RepositoryTypeProxy  = "proxy"
)

var (
	// RepositoryFormats contains a list of all supported repository formats
	RepositoryFormats = []string{
		RepositoryFormatApt,
		RepositoryFormatBower,
		RepositoryFormatConan,
		RepositoryFormatDocker,
		RepositoryFormatGitLFS,
		RepositoryFormatGo,
		RepositoryFormatHelm,
		RepositoryFormatMaven2,
		RepositoryFormatNPM,
		RepositoryFormatNuget,
		RepositoryFormatP2,
		RepositoryFormatPyPi,
		RepositoryFormatRAW,
		RepositoryFormatRuby,
		RepositoryFormatYum,
	}

	// RepositoryTypes contains a list of all supported repository types
	RepositoryTypes = []string{
		RepositoryTypeGroup,
		RepositoryTypeHosted,
		RepositoryTypeProxy,
	}
)

// LegacyRepository ...
type LegacyRepository struct {
	Format          string  `json:"format"`
	Name            string  `json:"name"`
	Online          bool    `json:"online"`
	RoutingRuleName *string `json:"routingRuleName,omitempty"`
	Type            string  `json:"type"`

	// Apt data
	*Apt        `json:"apt,omitempty"`
	*AptSigning `json:"aptSigning,omitempty"`

	// Bower data
	*Bower `json:"bower,omitempty"`

	// Cleanup data
	*Cleanup `json:"cleanup,omitempty"`

	// Docker data
	*Docker      `json:"docker,omitempty"`
	*DockerProxy `json:"dockerProxy,omitempty"`

	// Group data
	*Group `json:"group,omitempty"`

	// HTTPClient
	*HTTPClient `json:"httpClient,omitempty"`

	// Maven data
	*Maven `json:"maven,omitempty"`

	// Cache data for proxy Repository
	*NegativeCache `json:"negativeCache,omitempty"`

	// Nuget Proxy data
	*NugetProxy `json:"nugetProxy,omitempty"`

	// Proxy data
	*Proxy `json:"proxy,omitempty"`

	// Storage data
	Storage *HostedStorage `json:"storage,omitempty"`

	// Yum data
	*Yum `json:"yum,omitempty"`

	// Components
	*Component `json:"component,omitempty"`
}
