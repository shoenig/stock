package tool

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func findFile() string {
	// todo: XDG
	return filepath.Join(os.Getenv("HOME"), ".config/stock/stock.json")
}

func Load(paths []string) (map[string][]string, error) {
	return load(paths[0])
}

func load(path string) (map[string][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	m := make(map[string][]string)

	if decErr := json.NewDecoder(f).Decode(&m); decErr != nil {
		return nil, decErr
	}

	return m, nil
}
