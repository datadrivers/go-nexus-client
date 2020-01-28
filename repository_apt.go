package client

// RepositoryApt contains the data of an Apt Repository
type RepositoryApt struct {
	Distribution string `json:"distribution"`
}

// RepositoryAptSigning contains values for Apt signing
type RepositoryAptSigning struct {
	Keypair    string `json:"keypair"`
	Passphrase string `json:"passphrase"`
}

func (c client) RepositoryAptCreate(repo Repository, repoType string) error {
	return c.RepositoryCreate(repo, "apt", repoType)
}

func (c client) RepositoryAptRead(id string, repoType string) (*Repository, error) {
	return c.RepositoryRead(id, "apt", repoType)
}

func (c client) RepositoryAptUpdate(id string, repo Repository, repoType string) error {
	return c.RepositoryUpdate(id, repo, "apt", repoType)
}

func (c client) RepositoryAptDelete(id string) error {
	return c.RepositoryDelete(id)
}
