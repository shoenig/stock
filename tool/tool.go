package tool

import (
	"fmt"
	"os"

	yahoofinance "gophers.dev/pkgs/yahoo-finance"
)

type CLI struct {
	yf yahoofinance.Client
}

func New(yf yahoofinance.Client) *CLI {
	return &CLI{yf: yf}
}

func (cli *CLI) Run(tickers []string) error {
	if len(tickers) == 0 {
		return cli.runFile(findFile())
	}

	for _, ticker := range tickers {
		if err := cli.run(ticker, ""); err != nil {
			return err
		}
	}
	return nil
}

func (cli *CLI) runFile(filename string) error {
	m, loadErr := Load([]string{filename})
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
