package schema

// Script describe a groovy script object that can be run on the nexus server
type Script struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Type    string `json:"type"`
}
