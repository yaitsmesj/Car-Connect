package scheduler

import (
	"time"

	"github.com/yaitsmesj/Car-Connect/api-service/messanger"
)

// Schedule starts 2 Minute Ticker for Sending Event
func Schedule() chan bool {
	ticker := time.NewTicker(2 * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				messanger.SendMessage()
			}
		}
	}()
	return done
}
