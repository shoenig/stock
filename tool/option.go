package tool

import (
	"os"
	"path/filepath"

	"github.com/shoenig/yahoo-finance"
)

type Option func(c *CLI)

func WithClient(yf yahoofinance.Client) Option {
	return func(c *CLI) {
		c.yf = yf
	}
}

func WithDefaults(defaults string) Option {
	return func(c *CLI) {
		c.defaults = defaults
	}
}

func WithFile(name string) Option {
	return func(c *CLI) {
		p := filepath.Join(os.Getenv("HOME"), ".config", name)
		if _, err := os.Stat(p); err == nil {
			c.filepath = p
		}
	}
}
