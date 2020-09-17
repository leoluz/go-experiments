package main

import (
	"context"
	"time"

	"github.com/leoluz/go-experiments/bestpractices/context/heavy"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 2)
		defer cancel()
	}()
	heavy.Process(ctx)
}
