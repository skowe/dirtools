package main

import (
	"log"
	"sync"
	"time"

	"github.com/skowe/dirtools/monitor"
)

const relative = "main/targetFolder"

func logEvent(ch <-chan *monitor.Message, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	wg.Add(1)
	for m := range ch {
		log.Printf("EVENT: detected new file %s in folder %s\n", m.FileName, m.Path)
	}
}
func main() {
	sigChan := make(chan struct{}, 1)
	mon, err := monitor.InitMonitor(relative, 2)
	if err != nil {
		log.Panicln(err)
	}
	wg := &sync.WaitGroup{}
	log.Println(mon.Directory.Contents)
	go mon.Start(sigChan, wg)
	go logEvent(mon.InputCh, wg)
	for i := 0; i <= 10; i++ {
		sigChan <- struct{}{}
		time.Sleep(1 * time.Second)
	}
	close(sigChan)
	wg.Wait()
	log.Println(mon.Directory.Contents)
}
