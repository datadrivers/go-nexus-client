package repository

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/apt"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/bower"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/cocoapods"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/conan"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/conda"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/docker"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/gitlfs"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/golang"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/helm"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/legacy"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/maven"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/npm"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/nuget"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/p2"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/pypi"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/r"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/raw"
)

type RepositoryService struct {
	client *client.Client

	// API Services
	Apt       *apt.RepositoryAptService
	Bower     *bower.RepositoryBowerService
	Cocoapods *cocoapods.RepositoryCocoapodsService
	Conan     *conan.RepositoryConanService
	Conda     *conda.RepositoryCondaService
	Docker    *docker.RepositoryDockerService
	GitLfs    *gitlfs.RepositoryGitLfsService
	Go        *golang.RepositoryGoService
	Helm      *helm.RepositoryHelmService
	Maven     *maven.RepositoryMavenService
	Npm       *npm.RepositoryNpmService
	Nuget     *nuget.RepositoryNugetService
	P2        *p2.RepositoryP2Service
	Pypi      *pypi.RepositoryPypiService
	R         *r.RepositoryRService
	Raw       *raw.RepositoryRawService
	Legacy    *legacy.RepositoryLegacyService
}

func NewRepositoryService(c *client.Client) *RepositoryService {
	return &RepositoryService{
		client: c,

		Apt:       apt.NewRepositoryAptService(c),
		Bower:     bower.NewRepositoryBowerService(c),
		Cocoapods: cocoapods.NewRepositoryCocoapodsService(c),
		Conan:     conan.NewRepositoryConanService(c),
		Conda:     conda.NewRepositoryCondaService(c),
		Docker:    docker.NewRepositoryDockerService(c),
		GitLfs:    gitlfs.NewRepositoryGitLfsService(c),
		Go:        golang.NewRepositoryGoService(c),
		Helm:      helm.NewRepositoryHelmService(c),
		Maven:     maven.NewRepositoryMavenService(c),
		Npm:       npm.NewRepositoryNpmService(c),
		Nuget:     nuget.NewRepositoryNugetService(c),
		P2:        p2.NewRepositoryP2Service(c),
		Pypi:      pypi.NewRepositoryPypiService(c),
		R:         r.NewRepositoryRService(c),
		Raw:       raw.NewRepositoryRawService(c),
		Legacy:    legacy.NewRepositoryLegacyService(c),
	}
}
