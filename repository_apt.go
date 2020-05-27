package client

// RepositoryApt contains data of hosted repositories of format Apt
type RepositoryApt struct {
	Distribution string `json:"distribution,omitempty"`
	Flat         bool   `json:"flat"`
}

// RepositoryAptSigning contains signing data of hosted repositores of format Apt
type RepositoryAptSigning struct {
	Keypair    string `json:"keypair"`
	Passphrase string `json:"passphrase"`
}
