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

func (cli *CLI) Run(tickers []string) {
	for _, ticker := range tickers {
		cli.run(ticker)
	}
}

func (cli *CLI) run(ticker string) {
	chart, err := cli.yf.Lookup(ticker)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	prefix := ""
	dPerc := chart.DeltaPerc()
	if dPerc >= 0 {
		prefix = "+"
	}

	fmt.Printf("%s %.2f %s%.2f (%s%.2f%%)\n", chart.Symbol(), chart.Price(), prefix, chart.Delta(), prefix, chart.DeltaPerc())
}
