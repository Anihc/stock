package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type StockData struct {
	Code string         `json:"code"`
	Date int            `json:"date"`
	Time int            `json:"time"`
	Snap [2]interface{} `json:"snap"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a stock code as an argument")
		return
	}

	// Get the stock code from the first command-line argument
	stockCode := os.Args[1]
	// Instantiate a new collector
	c := colly.NewCollector()

	// Define the URL to scrape
	url := "http://yunhq.sse.com.cn:32041/v1/sh1/snap/" + stockCode + "?callback=jQuery112406539262313764276_1680657405102&select=prev_close%2Cname&_=1680657405105"

	// Define a slice to store the stock data
	var data StockData

	// Parse the JSON response
	c.OnResponse(func(r *colly.Response) {

		// Remove the callback function from the response body
		body := r.Body[42 : len(r.Body)-1]
		// fmt.Printf(string(body))
		// // Unmarshal the JSON data into the data slice
		if err := json.Unmarshal([]byte(body), &data); err != nil {
			log.Fatal(err)
		}
	})

	// Handle errors
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with error:", err)
	})

	// Start the scraping process
	c.Visit(url)

	// Print the stock data
	fmt.Printf("上交所-股票(%+v)：%+v, 价格：%+v\n", stockCode, data.Snap[1:], data.Snap[0:1])

	// fmt.Println(data)
}
