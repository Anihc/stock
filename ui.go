package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

type MarketData struct {
	Datetime         string `json:"datetime"`
	Code             string `json:"code"`
	MarketDataDetail struct {
		Code         string  `json:"code"`
		Name         string  `json:"name"`
		Close        string  `json:"close"`
		Open         string  `json:"open"`
		Now          string  `json:"now"`
		High         string  `json:"high"`
		Low          string  `json:"low"`
		Volume       int     `json:"volume"`
		Amount       float64 `json:"amount"`
		Delta        string  `json:"delta"`
		DeltaPercent string  `json:"deltaPercent"`
		LastVolume   int     `json:"lastVolume"`
		MarketTime   string  `json:"marketTime"`
		Sellbuy5     []struct {
			Price  string `json:"price"`
			Volume int    `json:"volume"`
		} `json:"sellbuy5"`
		Sellbuy1 []struct {
			Price  string `json:"price"`
			Volume int    `json:"volume"`
		} `json:"sellbuy1"`
		Picupdata         [][]interface{} `json:"picupdata"`
		Picdowndata       [][]interface{} `json:"picdowndata"`
		Picavgprice       [][]interface{} `json:"picavgprice"`
		GroupID           int             `json:"groupId"`
		VolumeAhT         int             `json:"volumeAhT"`
		AmountAhT         float64         `json:"amountAhT"`
		TradingPhaseCode1 string          `json:"tradingPhaseCode1"`
		TradingPhaseCode2 string          `json:"tradingPhaseCode2"`
		IsCDR             bool            `json:"isCDR"`
		IsNoProfit        int             `json:"isNoProfit"`
		IsVoteDifferent   int             `json:"isVoteDifferent"`
		IsVIE             int             `json:"isVIE"`
		IsRegistration    int             `json:"isRegistration"`
		Change20PerLimit  bool            `json:"change20PerLimit"`
		IsDelisting       interface{}     `json:"isDelisting"`
	} `json:"data"`
	Message string `json:"message"`
}

func main() {
	// Check if the user provided a stock code argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a stock code as an argument")
		return
	}

	// Get the stock code from the first command-line argument
	stockCode := os.Args[1]

	// Print the stock code
	fmt.Println("Stock code:", stockCode)

	// Create a new Colly collector
	c := colly.NewCollector()

	// Define a variable to store the parsed data
	var marketData MarketData

	// Visit the URL and parse the response JSON data
	c.OnResponse(func(r *colly.Response) {
		//fmt.Printf("%+v\n", string(r.Body))
		if err := json.Unmarshal(r.Body, &marketData); err != nil {
			log.Fatal(err)
		}
	})

	// Send a GET request to the URL
	if err := c.Visit("http://www.szse.cn/api/market/ssjjhq/getTimeData?random=0.5766232978431931&marketId=1&code=" + stockCode); err != nil {
		log.Fatal(err)
	}

	// Print the parsed data
	fmt.Printf("深交所-股票：%+v (%+v) ，当前价格：%+v\n", marketData.MarketDataDetail.Code, marketData.MarketDataDetail.Name, marketData.MarketDataDetail.Now)
}
