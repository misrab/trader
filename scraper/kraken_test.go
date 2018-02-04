package scraper

import (
	"testing"

	"fmt"
)

func TestKrakenTime(t *testing.T) {
	println("Testing time...")

	kraken := NewKraken("", "")
	time, err := kraken.Time()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Time is %v\n", time)
}
