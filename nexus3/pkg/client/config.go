package client

// Config is the configuration structure used to instantiate the Nexus client
type Config struct {
	URL                   string  `json:"url"`
	Username              string  `json:"username"`
	Password              string  `json:"password"`
	Insecure              bool    `json:"insecure"`
	Timeout               *int    `json:"timeout,omitempty"`
	ClientCertificatePath *string `json:"client_cert_path,omitempty"`
	ClientKeyPath         *string `json:"client_key_path,omitempty"`
	RootCAPath            *string `json:"root_ca_path,omitempty"`
}
