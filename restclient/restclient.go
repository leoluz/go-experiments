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

type ErrorMsg struct {
	StatusCode int
	Message    string
	DocUrl     string `json:"documentation_url"`
}

func (errorMsg ErrorMsg) Error() string {
	return fmt.Sprintf("Error code: %v %v: %v\n", errorMsg.StatusCode, errorMsg.Message, errorMsg.DocUrl)
}

type GithubRepos []GithubRepo

func main() {
	printUser()
	printRepos()
}

func get(url string) (response *http.Response, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	response, err = client.Do(req)
	if err != nil {
		return
	}

	if response.StatusCode >= 400 {
		errMsg := ErrorMsg{}
		json.NewDecoder(response.Body).Decode(&errMsg)
		errMsg.StatusCode = response.StatusCode
		err = errMsg
	}
	return
}

// Parse json response using Unmarshal
func printUser() {
	resp, err := get("https://api.github.com/users/leoluz")
	if err != nil {
		fmt.Println("Error when getting user info:", err)
		return
	}

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
	resp, err := get("https://api.github.com/users/leoluz/repos")
	if err != nil {
		fmt.Println("Error when getting repos:", err)
		return
	}

	repos := make(GithubRepos, 0)

	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err == nil {
	} else {
		fmt.Println("Error:", err)
	}
	fmt.Println("User repos:")
	for _, repo := range repos {
		fmt.Println("*", repo.Name, "-", repo.Description)
	}
}
