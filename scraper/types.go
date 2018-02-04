package scraper

import ()

type KrakenResponse struct {
	Error  []string    `json:json"error"`
	Result interface{} `json:"result"`
}

type TimeResponse struct {
	Unixtime int64
	Rfc1123  string
}
