Yahoo Finance API Portfolio Generator
========

The following code wil generate a "portfolio" (e.g. "folio.smith.json") for an array of given stocks (e.g. "YHOO", "AAPL", and "GOOG").

```go

package main

import "github.com/openwonk/yafin"

func main() {
	symbols := []string{"YHOO", "AAPL", "GOOG"}
	name := "Smith"
	
	yafin.CreatePortfolio(symbols, name)
}

```