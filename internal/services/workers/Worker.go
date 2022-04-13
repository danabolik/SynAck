package workers

import (
	"SynAck/internal/delivery"
	"SynAck/internal/services/decorators"
	"SynAck/internal/services/producers"
	"sync"
)

type Worker struct {
	Decorator decorators.DialerDecorator
	Delivery  delivery.Delivery
	Producer  producers.Producer
}

func (w Worker) ScanPorts(addr string, grt int) []int {
	tcp := w.Delivery.GetTcpNetwork()
	wg := sync.WaitGroup{}

	psChan := make(chan int, w.Producer.GetCountPorts())
	go w.Producer.WritePsToChan(psChan)
	var m sync.Mutex

	var result []int
	for i := 0; i < grt; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range psChan {
				dial := w.Decorator.DialPort(tcp, addr, p)
				if dial != 0 {
					m.Lock()
					result = append(result, dial)
					m.Unlock()
				}
			}
		}()
	}
	wg.Wait()

	return result
}
