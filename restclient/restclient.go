package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GithubRepo struct {
	Name        string
	Description string
}

type User struct {
	Login string
	Id    int32
	Name  string
}

type GithubRepos []GithubRepo

func main() {
	printUser()
	printRepos()
}

// Parse json response using Unmarshal
func printUser() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/users/leoluz", nil)
	resp, _ := client.Do(req)

	user := User{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &user)
	fmt.Println("User:")
	fmt.Printf("* %+v(%v): %v\n", user.Login, user.Id, user.Name)
}

// Parse json response using Decoder
func printRepos() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/users/leoluz/repos", nil)
	resp, _ := client.Do(req)
	repos := make(GithubRepos, 0)

	err := json.NewDecoder(resp.Body).Decode(&repos)
	if err == nil {
	} else {
		fmt.Println("Error:", err)
	}
	fmt.Println("User repos:")
	for _, repo := range repos {
		fmt.Println("*", repo.Name, "-", repo.Description)
	}
}
