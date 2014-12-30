package main

import (
	// "fmt"
	y "github.com/openwonk/yafin"
)

func main() {

	symbols := []string{"YHOO", "AAPL", "GOOG"}
	y.CreatePortfolio(symbols, "Davis")

}
