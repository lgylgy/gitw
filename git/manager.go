package git

func NewManager(repos *Repositories) *Manager {
	return &Manager{
		repos: repos,
		actions: []*Action{
			{
				Label: "fetch -> reset --hard upstream",
				Process: func(dir string, msg chan *Result) {
					output, err := fetch(dir, "upstream")
					msg <- &Result{
						Output: output,
						Err:    err,
					}
					if err != nil {
						return
					}
					output, err = resetHard(dir, "upstream", "master")
					msg <- &Result{
						Output: output,
						Err:    err,
					}
				},
			},
			{
				Label: "push origin HEAD",
				Process: func(dir string, msg chan *Result) {
					output, err := push(dir, "origin", "HEAD")
					msg <- &Result{
						Output: output,
						Err:    err,
					}
				},
			},
		},
		current: 0,
	}
}

type Manager struct {
	repos   *Repositories
	actions []*Action
	current int
}

func (m *Manager) Select(index int) *Repository {
	repo := m.repos.get(index)
	if repo == nil {
		return nil
	}
	m.current = index
	return repo
}

func (m *Manager) ListRepos() []string {
	return m.repos.getNames()
}

func (m *Manager) ListActions() []string {
	result := []string{}
	for _, action := range m.actions {
		result = append(result, action.Label)
	}
	return result
}

func (m *Manager) GetAction(index int) *Action {
	if index >= len(m.ListActions()) {
		return nil
	}
	return m.actions[index]
}
