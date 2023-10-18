package monitor

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// Aggregates multiple monitors at once and allows simultanious management of each
type Agregator struct {
	Monitors   []*Monitor
	ScanSignal chan struct{}
	Stopper    chan struct{}
}

func New(dirs []string) *Agregator {
	agregator := &Agregator{
		Monitors:   make([]*Monitor, 0),
		ScanSignal: make(chan struct{}),
		Stopper:    make(chan struct{}, 1),
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

func (a *Agregator) Start(worker Worker, Wg *sync.WaitGroup) {
	defer func() {
		Wg.Done()
	}()
	Wg.Add(1)
	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick.C:
			for _, mon := range a.Monitors {
				mon.Scan()
				go worker.Work(mon.InputCh)
			}
		case <-a.Stopper:
			fmt.Println("Stopping...")
			tick.Stop()
			for _, mon := range a.Monitors {
				close(mon.InputCh)
			}
			return
		}
	}
}

func (a *Agregator) Stop() {
	fmt.Println("Press enter to exit the program")
	scan := bufio.NewReader(os.Stdin)
	scan.ReadLine()

	close(a.Stopper)
}
