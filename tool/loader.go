package tool

import (
	"encoding/json"
	"os"
)

func Load(paths []string) (map[string][]string, error) {
	return load(paths[0])
}

func load(path string) (map[string][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	m := make(map[string][]string)

	if err := json.NewDecoder(f).Decode(&m); err != nil {
		return nil, err
	}

	return m, nil
}
