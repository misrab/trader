package scraper


import (
  "testing"
  //"log"
)

func TestKrakenTimeAndTicker(t *testing.T) {
  // create client
  kraken := NewKraken("", "")

  // test the time
  time, err := kraken.Time()
  if err != nil { t.Fatalf("%v\n", err) }
  // log.Printf("%v\n", time)
  if time.Unixtime == 0 {
    t.Fatal("Kraken unix time is 0")
  }

  // test eth-usd ticker
  ticker, err := kraken.TickerXETHZUSD()
  if err != nil { t.Fatalf("%v\n", err) }
  if len(ticker.Ask) == 0 {
    t.Fatal("Ticker for XETHUSD has no asks")
  }
  // log.Printf("%v\n", ticker)
}
