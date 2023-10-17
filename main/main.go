package main

import (
	"log"
	"sync"

	"github.com/skowe/dirtools/monitor"
)

const (
	relative1 = "main/targetFolder"
	relative2 = "main/target2"
)

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

}
