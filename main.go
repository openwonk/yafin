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

func RequestData(symbol string) {
	resp, err := http.Get("http://ichart.finance.yahoo.com/table.csv?s=" + symbol)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	// ioutil.TempDir(dir, prefix)
	check(ioutil.WriteFile("data."+strings.ToLower(symbol)+".csv", body, 0644))
}

func JSONize(symbol string) Stock {
	data, err := os.Open("data." + strings.ToLower(symbol) + ".csv")
	check(err)
	defer data.Close()

	reader := csv.NewReader(data)
	reader.FieldsPerRecord = -1 // see the Reader struct information below

	raw, err := reader.ReadAll()
	check(err)

	var oneSession TradingSession
	var multiSessions []TradingSession

	// TODO: skip first row (headers)... maybe if key > 0?
	for _, row := range raw {
		// fmt.Printf("%s %s\n", row[0], row[1])
		oneSession.Date = row[0]
		oneSession.Open, _ = strconv.ParseFloat(row[1], 64)  // strconv.Atoi(row[1])
		oneSession.High, _ = strconv.ParseFloat(row[2], 64)  // strconv.Atoi(row[2])
		oneSession.Low, _ = strconv.ParseFloat(row[3], 64)   // strconv.Atoi(row[3])
		oneSession.Close, _ = strconv.ParseFloat(row[4], 64) // strconv.Atoi(row[4])
		oneSession.Volume, _ = strconv.Atoi(row[5])
		oneSession.AdjClose, _ = strconv.ParseFloat(row[6], 64) // strconv.Atoi(row[6])
		multiSessions = append(multiSessions, oneSession)

	}

	// StockHistory := map[string][]TradingSession{symbol: multiSessions}
	// MarketHistory := map[string][]TradingSession{symbol: multiSessions}
	singleStock := Stock{Name: strings.ToUpper(symbol), History: multiSessions}

	jsonData, err := json.Marshal(singleStock)
	check(err)

	// fmt.Println(string(jsonData))

	jsonFile, err := os.Create("data." + symbol + ".json")
	check(err)

	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()

	return Stock

}

func GroupStocks(symbols []string) {
	var m Market

	for _, v := range symbols {
		RequestData(v)
		s := JSONize(v)
		m = append(m, s)
	}
}

type Market []Stock

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
