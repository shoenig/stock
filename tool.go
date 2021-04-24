package main

import (
	"fmt"
)

func getTickers(args []string) ([]string, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("supply at least one ticker symbol")
	}

	return args[1:], nil
}
