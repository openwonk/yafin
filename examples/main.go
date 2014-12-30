package main

import (
// yafin "github.com/openwonk/yafin.go"
// "encoding/csv"
// "encoding/json"
// "fmt"
// "os"
// "strconv"
// "strings"
)

func main() {
	// yafin.RequestData("YHOO")
	// yafin.JSONize("YHOO")

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
