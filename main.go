package yafin

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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

	// fmt.Println(string(data))

	// TODO: iterate through CSV and instert into JSON struct

	reader := csv.NewReader(data)

	raw, err := reader.ReadAll()
	check(err)

	for _, row := range raw {
		fmt.Printf(">> %s", row[0])
	}

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
