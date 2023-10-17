package monitor

import (
	"log"
	"sync"

	"github.com/skowe/dirtools/dirwrapper"
)

// Monitor sends a Message down an input channel
// Pass the input channel to share the data from monitor to your functions
type Monitor struct {
	Directory *dirwrapper.Directory
	InputCh   chan *Message
}

// Messages are sent via the Monitors Input channel
type Message struct {
	Path     string
	FileName string
}

func InitMonitor(directory string, bufferSize int) (*Monitor, error) {
	res := &Monitor{}
	dir, err := dirwrapper.Open(directory)

	if err != nil {
		return res, err
	}

	res.Directory = dir
	res.InputCh = make(chan *Message, bufferSize)
	return res, nil
}

// Once the signal channel closes it will close the message input channel
func (m *Monitor) Start(signal <-chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		close(m.InputCh)
		wg.Done()
	}()

	wg.Add(1)
	for range signal {
		scan(m)
	}

}

func scan(m *Monitor) {
	update, err := m.Directory.CheckForUpdate()
	if err != nil {
		log.Fatalln("FATAL ERROR: error when starting a monitor process:", err)
	}
	for _, filename := range update {
		message := &Message{
			Path:     m.Directory.Dir,
			FileName: filename,
		}
		m.InputCh <- message
	}
}
