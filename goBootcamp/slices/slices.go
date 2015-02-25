package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	maxlen := 0
	for _, v := range names {
		if lenth := len(v); lenth > maxlen {
			maxlen = lenth
		}
	}
	groups := make([][]string, maxlen)

	for _, name := range names {
		idx := len(name) - 1
		groups[idx] = append(groups[idx], name)
	}

	fmt.Println(groups)

}
