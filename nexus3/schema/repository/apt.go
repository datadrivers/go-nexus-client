package repository

// Apt contains data of hosted repositories of format Apt
type Apt struct {
	Distribution string `json:"distribution,omitempty"`
	Flat         bool   `json:"flat"`
}

// AptSigning contains signing data of hosted repositores of format Apt
type AptSigning struct {
	Keypair    string `json:"keypair"`
	Passphrase string `json:"passphrase"`
}
