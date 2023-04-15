package tool

import (
	"encoding/json"
	"strings"
)

func Load(groups string) (map[string][]string, error) {
	m := make(map[string][]string)
	r := strings.NewReader(groups)
	if decErr := json.NewDecoder(r).Decode(&m); decErr != nil {
		return nil, decErr
	}
	return m, nil
}
