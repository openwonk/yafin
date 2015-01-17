package yafin

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// TODO:
// - Parse and restructure dates for better searchability
// - Generate folders (if not exists) for "stock" and "folio" data
// - Refactor code to be leaner

func RequestData(symbol string) {
	resp, err := http.Get("http://ichart.finance.yahoo.com/table.csv?s=" + symbol)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	// ioutil.TempDir(dir, prefix)
	check(ioutil.WriteFile("stock."+strings.ToLower(symbol)+".csv", body, 0644))
}

func JsonCSV(symbol string) Stock {
	data, err := os.Open("stock." + strings.ToLower(symbol) + ".csv")
	check(err)
	defer data.Close()

	reader := csv.NewReader(data)
	reader.FieldsPerRecord = -1 // see the Reader struct information below

	raw, err := reader.ReadAll()
	check(err)

	var oneSession TradingSession
	var multiSessions []TradingSession

	for k, row := range raw {
		if k > 0 {
			oneSession.Date = row[0]
			oneSession.Open, _ = strconv.ParseFloat(row[1], 64)  // strconv.Atoi(row[1])
			oneSession.High, _ = strconv.ParseFloat(row[2], 64)  // strconv.Atoi(row[2])
			oneSession.Low, _ = strconv.ParseFloat(row[3], 64)   // strconv.Atoi(row[3])
			oneSession.Close, _ = strconv.ParseFloat(row[4], 64) // strconv.Atoi(row[4])
			oneSession.Volume, _ = strconv.Atoi(row[5])
			oneSession.AdjClose, _ = strconv.ParseFloat(row[6], 64) // strconv.Atoi(row[6])
			multiSessions = append(multiSessions, oneSession)
		}
	}

	singleStock := Stock{Name: strings.ToUpper(symbol), History: multiSessions}

	jsonData, err := json.Marshal(singleStock)
	check(err)

	jsonFile, err := os.Create("data." + strings.ToLower(symbol) + ".json")
	check(err)

	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

	return singleStock

}

func CreatePortfolio(symbols []string, name string) {
	var p []Stock

	for _, v := range symbols {
		RequestData(v)
		s := JsonCSV(v)
		p = append(p, s)
	}

	folio := Portfolio{Name: name, Stocks: p}

	jsonData, err := json.Marshal(folio)
	check(err)

	jsonFile, err := os.Create("folio." + strings.ToLower(string(name[:5])) + ".json")
	check(err)

	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}

// func JsonFolio(s Portfolio) {
// 	jsonData, err := json.Marshal(s)
// 	check(err)

// 	jsonFile, err := os.Create("folio." + strings.ToLower(string(s.Name[:5])) + ".json")
// 	check(err)

// 	defer jsonFile.Close()

// 	jsonFile.Write(jsonData)
// 	jsonFile.Close()
// }

type Portfolio struct {
	Name   string
	Stocks []Stock
}

type Stock struct {
	Name    string
	History []TradingSession
}

type TradingSession struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int
	AdjClose float64
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
