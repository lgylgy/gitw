package git

func NewManager(repos *Repositories) *Manager {
	return &Manager{
		repos: repos,
		actions: []*Action{
			{
				"fetch -> reset --hard upstream",
			},
			{
				"push origin HEAD",
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
