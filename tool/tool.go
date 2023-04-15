package tool

import (
	"fmt"
	"io"
	"os"

	yahoofinance "github.com/shoenig/yahoo-finance"
)

type CLI struct {
	filepath string
	defaults string
	yf       yahoofinance.Client
}

func New(opts ...Option) *CLI {
	cli := new(CLI)
	for _, opt := range opts {
		opt(cli)
	}
	return cli
}

func (cli *CLI) Run(tickers []string) error {
	if len(tickers) > 0 {
		for _, ticker := range tickers {
			if err := cli.run(ticker, ""); err != nil {
				return err
			}
		}
		return nil
	}

	var groups = cli.defaults
	if rc := cli.open(); rc != nil {
		custom, err := io.ReadAll(rc)
		if err != nil {
			return err
		}
		groups = string(custom)
	}

	return cli.runGroups(groups)
}

func (cli *CLI) open() io.ReadCloser {
	if cli.filepath == "" {
		return nil
	}

	f, err := os.Open(cli.filepath)
	if err != nil {
		return nil
	}

	return f
}

func (cli *CLI) runGroups(groups string) error {
	m, loadErr := Load(groups)
	if loadErr != nil {
		return loadErr
	}

	for label, list := range m {
		fmt.Printf("[%s]\n", label)
		for _, ticker := range list {
			if err := cli.run(ticker, "\t"); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to lookup %q: %v\n", ticker, err)
			}
		}
	}

	return nil
}

func (cli *CLI) run(ticker, ident string) error {
	chart, err := cli.yf.Lookup(ticker)
	if err != nil {
		return err
	}

	prefix := ""
	dPerc := chart.DeltaPerc()
	if dPerc >= 0 {
		prefix = "+"
	}

	fmt.Printf("%s%s %.2f %s%.2f (%s%.2f%%)\n", ident, chart.Symbol(), chart.Price(), prefix, chart.Delta(), prefix, chart.DeltaPerc())
	return nil
}
