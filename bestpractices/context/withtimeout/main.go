package main

import (
	"context"
	"time"

	"github.com/leoluz/go-experiments/bestpractices/context/heavy"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	heavy.Process(ctx)
}
