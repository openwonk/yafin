Yahoo Finance API Portfolio Generator
========

The following code generates a "portfolio" JSON file for an array of given stocks (e.g. "YHOO", "AAPL", "GOOG").

```go

package main

import "github.com/openwonk/yafin"

func main() {
	symbols := []string{"YHOO", "AAPL", "GOOG"}
	
	yafin.CreatePortfolio(symbols, "Davis")
}

```