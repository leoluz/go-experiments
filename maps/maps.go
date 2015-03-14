package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordsCount := map[string]int{}
	for _, word := range strings.Fields(s) {
		elem, ok := wordsCount[word]
		if ok {
			wordsCount[word] = elem + 1
		} else {
			wordsCount[word] = 1
		}
	}
	return wordsCount
}

func main() {
	wc.Test(WordCount)
}
