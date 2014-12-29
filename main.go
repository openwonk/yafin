package yafin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// func main() {
// 	// RequestData("YHOO")
// 	JSONize("YHOO")
// }

func RequestData(symbol string) {
	resp, err := http.Get("http://ichart.finance.yahoo.com/table.csv?s=" + symbol)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	// ioutil.TempDir(dir, prefix)
	check(ioutil.WriteFile("data."+strings.ToLower(symbol)+".csv", body, 0644))
}

func JSONize(symbol string) {
	filepath := "data." + strings.ToLower(symbol) + ".csv"

	data, err := ioutil.ReadFile(filepath)
	check(err)

	fmt.Println(string(data))

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
