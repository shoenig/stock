package main

import (
	"fmt"
	"os"

	"gophers.dev/cmds/stock/tool"
	yahoofinance "gophers.dev/pkgs/yahoo-finance"
)

func getTickers(args []string) []string {
	return args[1:]
}

func main() {
	yf := yahoofinance.New(nil)
	tickers := getTickers(os.Args)
	cli := tool.New(yf)
	if err := cli.Run(tickers); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to run: %s", err)
	}
}
