package git

import (
	"encoding/json"
	"fmt"
	"os"
)

func NewRepositories() *Repositories {
	return &Repositories{
		list: []*Repository{},
	}
}

type Repositories struct {
	list []*Repository
}

func (r *Repositories) Load(config string) error {
	configFile, err := os.Open(config)
	if err != nil {
		return err
	}
	defer configFile.Close()

	// read config file
	list := []*Repository{}
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&list)
	if err != nil {
		return err
	}

	r.list = list
	if len(r.list) == 0 {
		return fmt.Errorf("no repository found")
	}
	return nil
}

func (r *Repositories) get(index int) *Repository {
	return r.list[index]
}

func (r *Repositories) getNames() []string {
	result := []string{}
	for _, item := range r.list {
		result = append(result, item.Name)
	}
	return result
}
