package scraper

import (
  "log"
  "time"
)


type Poller struct {
  target string
  interval time.Duration
}

func NewPoller(target string, interval time.Duration) *Poller {
  return &Poller {
    target: target,
    interval: interval,
  }
}

func (p *Poller) Poll(quit chan struct{}) {
  for {
    select {
      case <-quit: {
        log.Printf("Quitting poller on %s at interval %f\n", p.target, p.interval.Seconds())
      }
    }
  }
}
