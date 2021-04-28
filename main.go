package main

import (
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

	//fmt.Println("first trade:", chart.FirstTrade())
	//
	//fmt.Println("market time:", chart.MarketTime())
	//
	//fmt.Println("symbol:", chart.Symbol())
	//
	//fmt.Println("price:", chart.Price())
	//
	//fmt.Println("previous:", chart.PrevClosePrice())
	//
	//fmt.Println("price delta:", chart.Delta())
	//
	//fmt.Println("done")

	prefix := ""
	dPerc := chart.DeltaPerc()
	if dPerc >= 0 {
		prefix = "+"
	}

	fmt.Printf("%s %.2f %s%.2f (%s%.2f%%)\n", chart.Symbol(), chart.Price(), prefix, chart.Delta(), prefix, chart.DeltaPerc())
}
