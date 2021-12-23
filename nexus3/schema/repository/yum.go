package repository

// Yum contains data of hosted repositories of format Yum
type Yum struct {
	RepodataDepth int    `json:"repodataDepth"`
	DeployPolicy  string `json:"deployPolicy"`
}
