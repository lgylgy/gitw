package git

type Repository struct {
	Name string
	Path string
}

func (r *Repository) GetCurrentBranch() (string, error) {
	text, err := getCurrentBranch(r.Path)
	if err != nil {
		return "", err
	}
	return text, nil
}

func (r *Repository) GetRemotes() (string, error) {
	text, err := getRemotes(r.Path)
	if err != nil {
		return "", err
	}
	return text, nil
}
