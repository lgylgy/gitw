package git

type Repository struct {
	Name string
	Path string
}

func (r *Repository) GetCurrentBranch() (string, error) {
	branch, err := getCurrentBranch(r.Path)
	if err != nil {
		return "", err
	}
	return branch, nil
}

func (r *Repository) GetRemotes() ([]string, error) {
	remotes, err := getRemotes(r.Path)
	if err != nil {
		return nil, err
	}
	return remotes, nil
}
