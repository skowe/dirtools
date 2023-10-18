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

type Logger struct{}

func (l *Logger) Work(input any) {
	ch, ok := input.(chan *monitor.Message)

	if !ok {
		log.Println("FATAL: Failed after scan operation")
		return
	}
	log.Println(<-ch)
}
func main() {
	aggr := monitor.New([]string{relative1, relative2})
	Wg := &sync.WaitGroup{}
	go aggr.Start(&Logger{}, Wg)

	aggr.Stop()

	Wg.Wait()
}
