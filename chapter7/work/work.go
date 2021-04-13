package work

import "sync"

type Worker interface {
	Task()
}

// goroutine池
type Pool struct {
	worker chan Worker
	wg     sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		worker: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			//干活
			for w := range p.worker {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// run:提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.worker <- w
}

func (p *Pool) Shutdown() {
	close(p.worker)
	p.wg.Wait()
}
