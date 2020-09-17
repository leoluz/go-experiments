package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/leoluz/go-experiments/bestpractices/context/heavy"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		defer cancel()
		s := <-stop
		fmt.Printf("Received %s signal\n", s)
	}()

	heavy.Process(ctx)
}
