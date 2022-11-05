package pkgs

type Package struct {
	AuthorID    string `json:"author_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	GitURL      string `json:"git_url"`
}
