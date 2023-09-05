package apis

type Repo struct {
	RepoURL  string `json:"repoURL,omitempty" yaml:"repoURL,omitempty"`
	RepoPath string `json:"repoPath,omitempty" yaml:"repoPath,omitempty"`
}
