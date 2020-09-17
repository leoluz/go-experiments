package heavy

import (
	"context"
	"fmt"
	"time"
)

type HeavyProcess struct {
	stopped bool
}

func Process(ctx context.Context) {
	process := new(HeavyProcess)
	resultChan, doneChan := process.Start()

	for {
		select {
		case <-ctx.Done():
			process.stop()
		case result := <-resultChan:
			fmt.Println(result)
		case <-doneChan:
			fmt.Println("HeavyProcess is done!")
			return
		}
	}
}

func (p *HeavyProcess) Start() (<-chan string, <-chan struct{}) {
	result := make(chan string)
	done := make(chan struct{})
	go p.do(result, done)
	return result, done
}

func (p *HeavyProcess) do(result chan<- string, done chan<- struct{}) {
	defer close(result)
	defer close(done)

	for i := 0; i < 5; i++ {
		result <- fmt.Sprintf("Hi from %d", i)
		if p.stopped {
			done <- struct{}{}
			return
		}
		time.Sleep(time.Second * 1)
	}
	done <- struct{}{}
}

func (p *HeavyProcess) stop() {
	p.stopped = true
}
