package git

import (
	"encoding/json"
	"os"
)

func NewConfig() *Config {
	return &Config{
		Repositories: make(map[string]*Repository),
	}
}

type Repository struct {
	Name string
	Path string
}

type Config struct {
	Repositories map[string]*Repository
}

func (c *Config) UnmarshalJSON(data []byte) error {
	repos := []*Repository{}
	if err := json.Unmarshal(data, &repos); err != nil {
		return err
	}
	for _, repo := range repos {
		c.Repositories[repo.Name] = repo
	}
	return nil
}

func LoadConfiguration(file string) (*Config, error) {
	configFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	config := NewConfig()
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
