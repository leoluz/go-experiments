package main

import (
	"context"

	"github.com/leoluz/go-experiments/bestpractices/context/heavy"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()
	//stop := make(chan os.Signal, 1)
	//signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	//go func() {
	//defer cancel()
	//fmt.Printf("Received %s signal\n", <-stop)
	//}()

	heavy.Process(ctx)
}
