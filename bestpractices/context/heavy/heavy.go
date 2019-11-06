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
	resultChan := make(chan string)
	doneChan := make(chan struct{})
	defer close(resultChan)
	defer close(doneChan)

	process := new(HeavyProcess)
	go process.do(resultChan, doneChan)

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

func (p *HeavyProcess) do(result chan<- string, done chan<- struct{}) {
	for i := 0; i < 5; i++ {
		if p.stopped {
			done <- struct{}{}
			return
		}
		result <- fmt.Sprintf("Hi from %d", i)
		time.Sleep(time.Second * 1)
	}
	done <- struct{}{}
}

func (p *HeavyProcess) stop() {
	p.stopped = true
}
