package monitor

import (
	"log"
	"sync"
	"time"
)

// Aggregates multiple monitors at once and allows simultanious management of each
type Agregator struct {
	Monitors   []*Monitor
	ScanSignal chan struct{}
	Stopper    chan struct{}
	Wg         *sync.WaitGroup
}

func New(dirs []string) *Agregator {
	agregator := &Agregator{
		Monitors:   make([]*Monitor, 0),
		ScanSignal: make(chan struct{}),
		Stopper:    make(chan struct{}, 1),
		Wg:         &sync.WaitGroup{},
	}

	for _, dir := range dirs {
		mon, err := InitMonitor(dir, len(dirs)/2)
		if err != nil {
			log.Fatalf("FATAL: error %s\nINFO: failed to initialize directory monitor for %s", err, dir)
			return nil
		}

		agregator.Monitors = append(agregator.Monitors, mon)

	}

	return agregator
}

func (a *Agregator) Start(worker Worker) {
	tick := time.NewTicker(time.Second)

	for {
		select {
		case <-tick.C:
			for _, mon := range a.Monitors {
				mon.Scan()
				worker.Work(mon.InputCh)
			}
		case <-a.Stopper:
			tick.Stop()
			return
		}
	}
}

func (a *Agregator) Stop() {
	a.Stopper <- struct{}{}
}
