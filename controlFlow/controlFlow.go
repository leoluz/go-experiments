package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
	vowelValues  = map[string]int{"a": 1, "e": 1, "i": 2, "o": 3, "u": 4}
)

func coinsForUser(user string) (totalUser int) {
	for vowel, value := range vowelValues {
		count := strings.Count(strings.ToLower(user), vowel)
		totalUser += count * value
	}
	return
}

func main() {
	for _, user := range users {
		totalUser := coinsForUser(user)
		if totalUser > 10 {
			totalUser = 10
		}
		distribution[user] = totalUser
		coins -= totalUser
	}
	fmt.Println(distribution)
	fmt.Println("Coins left: ", coins)
}
