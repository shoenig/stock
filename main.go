package main

import (
	"fmt"
	"os"

	"github.com/shoenig/stock/tool"
	"github.com/shoenig/yahoo-finance"
)

func main() {
	cli := tool.New(
		tool.WithClient(yahoofinance.New(nil)),
		tool.WithFile("stock.json"),
		tool.WithDefaults(DefaultJSON),
	)
	if err := cli.Run(os.Args[1:]); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to run: %s\n", err)

	}
}
