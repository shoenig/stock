package main

import (
	"fmt"
	"os"

	"gophers.dev/cmds/stock/tool"
	yahoofinance "gophers.dev/pkgs/yahoo-finance"
)

func getTickers(args []string) ([]string, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("supply at least one ticker symbol")
	}
	return args[1:], nil
}

func main() {
	yf := yahoofinance.New(nil)

	tickers, err := getTickers(os.Args)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	cli := tool.New(yf)
	cli.Run(tickers)
}
