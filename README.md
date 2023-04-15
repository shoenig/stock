# stock

Use `stock` to see current stock ticker pricing

# Project Overview

Module `github.com/shoenig/stock` provides a command-line utility for reporting
current stock ticker pricing.

# Getting Started

The `stock` command can be installed by running
```bash
go install github.com/shoenig/stock@latest
```

# Optional Configuration

If `~/stock/stock.json` exists, it can be used to display groups of
lists of ticker pricing data.

```
$ cat ~/.config/stock/stock.json
{
    "semis": ["amd", "intc", "tsm", "nvda", "xlnx"],
    "memes": ["gme", "amc", "bb", "pltr"],
    "cars": ["f", "tsla", "vwagy", "gm"],
    "indx": ["qqq", "spy", "^dji"]
}
```

# Example Usages

### Single ticker

```bash
$ stock amd
AMD 94.70 +1.39 (+1.49%)
```

### Multiple tickers

```bash
$ stock amd intc nvda tsm
AMD 94.70 +1.39 (+1.49%)
INTC 56.76 +0.75 (+1.34%)
NVDA 819.48 +11.00 (+1.36%)
TSM 118.91 +0.49 (+0.41%)
```

### With config file

```
$ stock
[cars]
	F 14.93 +0.02 (+0.13%)
	TSLA 678.90 +0.98 (+0.14%)
	VWAGY 33.31 -0.04 (-0.10%)
	GM 58.96 -0.15 (-0.25%)
[indx]
	QQQ 358.64 +4.07 (+1.15%)
	SPY 433.72 +3.29 (+0.76%)
	^DJI 34786.35 +152.85 (+0.44%)
[semis]
	AMD 94.70 +1.39 (+1.49%)
	INTC 56.76 +0.75 (+1.34%)
	TSM 118.91 +0.49 (+0.41%)
	NVDA 819.48 +11.00 (+1.36%)
	XLNX 144.56 +2.75 (+1.94%)
[memes]
	GME 202.83 -1.53 (-0.75%)
	AMC 51.96 -2.26 (-4.17%)
	BB 11.99 -0.18 (-1.48%)
	PLTR 24.44 -0.28 (-1.13%)
```

# Contributing

The `github.com/shoenig/stock` module is always improving with new features and
error corrections. For contributing bug fixes and new features please file an
issue.

# License

The `github.com/shoenig/stock` module is open source under the [BSD](LICENSE)
