package client

// RepositoryBower contains data of bower repositories
type RepositoryBower struct {
	RewritePackageUrls bool `json:"rewritePackageUrls"`
}

func (c client) RepositoryBowerCreate(repo Repository, repoType string) error {
	return c.RepositoryCreate(repo, "bower", repoType)
}

func (c client) RepositoryBowerRead(id string, repoType string) (*Repository, error) {
	return c.RepositoryRead(id)
}

func (c client) RepositoryBowerUpdate(id string, repo Repository, repoType string) error {
	return c.RepositoryUpdate(id, repo, "bower", repoType)
}

func (c client) RepositoryBowerDelete(id string) error {
	return c.RepositoryDelete(id)
}
