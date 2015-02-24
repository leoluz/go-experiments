package main

import "fmt"

func average(scores []float64) float64 {
	total := 0.0
	for _, v := range scores {
		total += v
	}
	return total / float64(len(scores))
}

func main() {
	scores := []float64{3, 2, 5, 10}
	fmt.Println(average(scores))
}
