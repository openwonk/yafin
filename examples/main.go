package main

import (
	// "fmt"
	yafin "github.com/openwonk/yafin"
)

func main() {
	// s := "hola mondo"
	// fmt.Println(string(s[:3]))
	symbols := []string{"YHOO", "AAPL", "GOOG"}
	yafin.CreatePortfolio(symbols, "Davis")
}
