package scraper

import (
  "testing"
  "time"
)


func TestPoller(t *testing.T) {
  p := NewPoller("http://www.google.com", 100*time.Millisecond)
  quit := make(chan struct{})
  go p.Poll(quit)
  time.Sleep(500*time.Millisecond)
  close(quit)
}
