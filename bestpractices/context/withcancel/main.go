package main

import (
	"context"

	"github.com/leoluz/go-experiments/bestpractices/context/heavy"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	heavy.Process(ctx)
}
