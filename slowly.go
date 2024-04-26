package slowly

import "time"

type TickerEntry struct {
	Limit  int
	Incr   chan int
	Ticker *time.Ticker
	Done   chan bool
}

func NewTicker(limit int, duration time.Duration) (entry TickerEntry) {
	//entry := TickerEntry{Limit: limit}
	entry.Limit = limit
	entry.Ticker = time.NewTicker(duration)
	entry.Done = make(chan bool)
	entry.Incr = make(chan int, limit)

	go entry.Monitor()
	return entry
}

func (entry *TickerEntry) ResetIncr() {
	for i := len(entry.Incr); i < entry.Limit; i++ {
		entry.Incr <- i
	}
}
func (entry *TickerEntry) Monitor() {
	for {
		select {
		case <-entry.Ticker.C:
			entry.ResetIncr()
		case <-entry.Done:
			entry.Ticker.Stop()
			return

		}
	}

}

func (entry *TickerEntry) Exec() {
	_ = <-entry.Incr
}

func (entry *TickerEntry) Stop() {
	entry.Done <- true
}
