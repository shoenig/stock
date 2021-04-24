package main

import (
	"encoding/json"
	"fmt"
	"os"

	yahoofinance "gophers.dev/pkgs/yahoo-finance"
)

func main() {
	yf := yahoofinance.New(nil)

	tickers, err := getTickers(os.Args)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	chart, err := yf.Lookup(tickers[0])
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	b, err := json.Marshal(chart)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
	fmt.Println("chart:", string(b))

	fmt.Println("first trade:", chart.FirstTrade())

	fmt.Println("market time:", chart.MarketTime())

	fmt.Println("done")
}
