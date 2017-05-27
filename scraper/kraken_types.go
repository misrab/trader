package scraper

import (

)


const (
	// XETHXXBT = "XETHXXBT"
	// XETHZCAD = "XETHZCAD"
	// XETHZEUR = "XETHZEUR"
	// XETHZGBP = "XETHZGBP"
	// XETHZJPY = "XETHZJPY"
	XETHZUSD = "XETHZUSD"
	// XLTCZCAD = "XLTCZCAD"
	// XLTCZEUR = "XLTCZEUR"
	// XLTCZUSD = "XLTCZUSD"
	// XXBTXLTC = "XXBTXLTC"
	// XXBTXNMC = "XXBTXNMC"
	// XXBTXXDG = "XXBTXXDG"
	// XXBTXXLM = "XXBTXXLM"
	// XXBTXXRP = "XXBTXXRP"
	// XXBTZCAD = "XXBTZCAD"
	// XXBTZEUR = "XXBTZEUR"
	// XXBTZGBP = "XXBTZGBP"
	// XXBTZJPY = "XXBTZJPY"
	// XXBTZUSD = "XXBTZUSD"
	// XXMRZUSD = "XXMRZUSD"
	// XXMRZEUR = "XXMRZEUR"
	// XXMRXXBT = "XXMRXXBT"
)

const (
	BUY    = "b"
	SELL   = "s"
	MARKET = "m"
	LIMIT  = "l"
)


// https://www.kraken.com/help/api

// KrakenResponse wraps the Kraken API JSON response
// from https://github.com/beldur/kraken-go-api-client/blob/master/types.go/
type KrakenResponse struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

// TimeResponse represents the server's time
type TimeResponse struct {
	// Unix timestamp
	Unixtime int64
	// RFC 1123 time format
	Rfc1123 string
}


// PairTickerInfo represents ticker information for a pair
type PairTickerInfo struct {
	// Ask array(<price>, <whole lot volume>, <lot volume>)
	Ask []string `json:"a"`
	// Bid array(<price>, <whole lot volume>, <lot volume>)
	Bid []string `json:"b"`
	// Last trade closed array(<price>, <lot volume>)
	Close []string `json:"c"`
	// Volume array(<today>, <last 24 hours>)
	Volume []string `json:"v"`
	// Volume weighted average price array(<today>, <last 24 hours>)
	VolumeAveragePrice []string `json:"p"`
	// Number of trades array(<today>, <last 24 hours>)
	Trades []int `json:"t"`
	// Low array(<today>, <last 24 hours>)
	Low []string `json:"l"`
	// High array(<today>, <last 24 hours>)
	High []string `json:"h"`
	// Today's opening price
	OpeningPrice float32 `json:"o,string"`
}


// TickerResponse includes the requested ticker pairs
type TickerResponse struct {
	XETHXXBT PairTickerInfo
	XETHZCAD PairTickerInfo
	XETHZEUR PairTickerInfo
	XETHZGBP PairTickerInfo
	XETHZJPY PairTickerInfo
	XETHZUSD PairTickerInfo
	XLTCZCAD PairTickerInfo
	XLTCZEUR PairTickerInfo
	XLTCZUSD PairTickerInfo
	XXBTXLTC PairTickerInfo
	XXBTXNMC PairTickerInfo
	XXBTXXDG PairTickerInfo
	XXBTXXLM PairTickerInfo
	XXBTXXRP PairTickerInfo
	XXBTZCAD PairTickerInfo
	XXBTZEUR PairTickerInfo
	XXBTZGBP PairTickerInfo
	XXBTZJPY PairTickerInfo
	XXBTZUSD PairTickerInfo
	XXMRZUSD PairTickerInfo
	XXMRZEUR PairTickerInfo
	XXMRXXBT PairTickerInfo
}
